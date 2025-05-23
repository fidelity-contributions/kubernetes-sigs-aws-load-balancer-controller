package gateway

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	elbv2gw "sigs.k8s.io/aws-load-balancer-controller/apis/gateway/v1beta1"
	gatewayclasseventhandlers "sigs.k8s.io/aws-load-balancer-controller/controllers/gateway/eventhandlers/gatewayclass"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/config"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/gateway/constants"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	reasonConfigNotFound  = "NotFound"
	messageConfigNotFound = "LoadBalancerConfiguration (%s,%s) not found"
)

// NewGatewayClassReconciler constructs a reconciler that responds to gateway class object changes
func NewGatewayClassReconciler(k8sClient client.Client, eventRecorder record.EventRecorder, controllerConfig config.ControllerConfig, enabledControllers sets.Set[string], logger logr.Logger) Reconciler {

	return &gatewayClassReconciler{
		k8sClient:                   k8sClient,
		eventRecorder:               eventRecorder,
		logger:                      logger,
		enabledControllers:          enabledControllers,
		workers:                     controllerConfig.GatewayClassMaxConcurrentReconciles,
		updateGwClassAcceptedFn:     updateGatewayClassAcceptedCondition,
		updateLastProcessedConfigFn: updateGatewayClassLastProcessedConfig,
		configResolverFn:            resolveLoadBalancerConfig,
	}
}

// gatewayClassReconciler reconciles Gateway Classes.
type gatewayClassReconciler struct {
	k8sClient          client.Client
	eventRecorder      record.EventRecorder
	logger             logr.Logger
	enabledControllers sets.Set[string]
	workers            int

	updateGwClassAcceptedFn     func(ctx context.Context, k8sClient client.Client, gwClass *gwv1.GatewayClass, status metav1.ConditionStatus, reason string, message string) error
	updateLastProcessedConfigFn func(ctx context.Context, k8sClient client.Client, gwClass *gwv1.GatewayClass, lbConf *elbv2gw.LoadBalancerConfiguration) error
	configResolverFn            func(ctx context.Context, k8sClient client.Client, reference *gwv1.ParametersReference) (*elbv2gw.LoadBalancerConfiguration, error)
}

func (r *gatewayClassReconciler) SetupWatches(_ context.Context, ctrl controller.Controller, mgr ctrl.Manager) error {

	gwClassEventChan := make(chan event.TypedGenericEvent[*gwv1.GatewayClass])
	lbEventHandler := gatewayclasseventhandlers.NewEnqueueRequestsForLoadBalancerConfigurationEvent(gwClassEventChan, r.k8sClient, r.eventRecorder, r.enabledControllers, r.logger)

	if err := ctrl.Watch(source.Kind(mgr.GetCache(), &gwv1.GatewayClass{}, &handler.TypedEnqueueRequestForObject[*gwv1.GatewayClass]{})); err != nil {
		return err
	}

	if err := ctrl.Watch(source.Channel(gwClassEventChan, &handler.TypedEnqueueRequestForObject[*gwv1.GatewayClass]{})); err != nil {
		return err
	}

	if err := ctrl.Watch(source.Kind(mgr.GetCache(), &elbv2gw.LoadBalancerConfiguration{}, lbEventHandler)); err != nil {
		return err
	}

	return nil
}

// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gatewayclasses,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gatewayclasses/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=gateway.networking.k8s.io,resources=gatewayclasses/finalizers,verbs=update
// +kubebuilder:rbac:groups=gateway.k8s.aws,resources=loadbalancerconfigurations,verbs=get;list;watch
func (r *gatewayClassReconciler) Reconcile(ctx context.Context, req reconcile.Request) (ctrl.Result, error) {
	err := r.reconcile(ctx, req)
	return runtime.HandleReconcileError(err, r.logger)
}

func (r *gatewayClassReconciler) reconcile(ctx context.Context, req reconcile.Request) error {
	gwClass := &gwv1.GatewayClass{}
	if err := r.k8sClient.Get(ctx, req.NamespacedName, gwClass); err != nil {
		return client.IgnoreNotFound(err)
	}

	r.logger.V(1).Info("Found updated gateway class", "class", gwClass)

	if !r.enabledControllers.Has(string(gwClass.Spec.ControllerName)) {
		return nil
	}

	var lbConf *elbv2gw.LoadBalancerConfiguration

	lbConf, err := r.configResolverFn(ctx, r.k8sClient, gwClass.Spec.ParametersRef)
	if err != nil {
		if client.IgnoreNotFound(err) == nil {
			// Ignoring error, because we want to highlight the bad configuration, not the failed status update.
			r.updateGwClassAcceptedFn(ctx, r.k8sClient, gwClass, metav1.ConditionFalse, reasonConfigNotFound, r.getNotFoundMessage(gwClass.Spec.ParametersRef))
		}
		return err
	}

	err = r.updateLastProcessedConfigFn(ctx, r.k8sClient, gwClass, lbConf)
	if err != nil {
		r.logger.Error(err, "Unable to update last processed annotation")
		return err
	}

	err = r.updateGwClassAcceptedFn(ctx, r.k8sClient, gwClass, metav1.ConditionTrue, string(gwv1.GatewayClassReasonAccepted), string(gwv1.GatewayClassReasonAccepted))
	if err != nil {
		r.logger.Error(err, "Unable to update condition")
		return err
	}

	return nil
}

func (r *gatewayClassReconciler) getNotFoundMessage(paramRef *gwv1.ParametersReference) string {
	var ns string
	if paramRef.Namespace == nil {
		ns = string(*paramRef.Namespace)
	}

	return fmt.Sprintf(messageConfigNotFound, paramRef.Name, ns)
}

func (r *gatewayClassReconciler) SetupWithManager(_ context.Context, mgr ctrl.Manager) (controller.Controller, error) {
	return controller.New(constants.GatewayClassController, mgr, controller.Options{
		MaxConcurrentReconciles: r.workers,
		Reconciler:              r,
	})

}

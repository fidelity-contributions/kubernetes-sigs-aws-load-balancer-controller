package routeutils

import (
	"k8s.io/apimachinery/pkg/types"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	"time"
)

type MockRule struct {
	RawRule     interface{}
	BackendRefs []Backend
}

func (m *MockRule) GetRawRouteRule() interface{} {
	return m.RawRule
}

func (m *MockRule) GetSectionName() *gwv1.SectionName {
	//TODO implement me
	panic("implement me")
}

func (m *MockRule) GetBackends() []Backend {
	return m.BackendRefs
}

var _ RouteRule = &MockRule{}

type MockRoute struct {
	Kind         RouteKind
	Name         string
	Namespace    string
	Hostnames    []string
	CreationTime time.Time
	Rules        []RouteRule
}

func (m *MockRoute) GetBackendRefs() []gwv1.BackendRef {
	//TODO implement me
	panic("implement me")
}

func (m *MockRoute) GetListenerRuleConfigs() []gwv1.LocalObjectReference {
	//TODO implement me
	panic("implement me")
}

func (m *MockRoute) GetRouteNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: m.Namespace,
		Name:      m.Name,
	}
}

func (m *MockRoute) GetRouteKind() RouteKind {
	return m.Kind
}

func (m *MockRoute) GetHostnames() []gwv1.Hostname {
	hostnames := make([]gwv1.Hostname, len(m.Hostnames))
	for i, h := range m.Hostnames {
		hostnames[i] = gwv1.Hostname(h)
	}
	return hostnames
}

func (m *MockRoute) GetParentRefs() []gwv1.ParentReference {
	//TODO implement me
	panic("implement me")
}

func (m *MockRoute) GetRawRoute() interface{} {
	//TODO implement me
	panic("implement me")
}

func (m *MockRoute) GetAttachedRules() []RouteRule {
	//TODO implement me
	return m.Rules
}

func (m *MockRoute) GetRouteGeneration() int64 {
	panic("implement me")
}

func (m *MockRoute) GetRouteCreateTimestamp() time.Time {
	return m.CreationTime
}

var _ RouteDescriptor = &MockRoute{}

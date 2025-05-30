// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/aws-load-balancer-controller/pkg/aws/services (interfaces: EC2)

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	gomock "github.com/golang/mock/gomock"
)

// MockEC2 is a mock of EC2 interface.
type MockEC2 struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MockRecorder
}

// MockEC2MockRecorder is the mock recorder for MockEC2.
type MockEC2MockRecorder struct {
	mock *MockEC2
}

// NewMockEC2 creates a new mock instance.
func NewMockEC2(ctrl *gomock.Controller) *MockEC2 {
	mock := &MockEC2{ctrl: ctrl}
	mock.recorder = &MockEC2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2) EXPECT() *MockEC2MockRecorder {
	return m.recorder
}

// AuthorizeSecurityGroupIngressWithContext mocks base method.
func (m *MockEC2) AuthorizeSecurityGroupIngressWithContext(arg0 context.Context, arg1 *ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizeSecurityGroupIngressWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.AuthorizeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeSecurityGroupIngressWithContext indicates an expected call of AuthorizeSecurityGroupIngressWithContext.
func (mr *MockEC2MockRecorder) AuthorizeSecurityGroupIngressWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeSecurityGroupIngressWithContext", reflect.TypeOf((*MockEC2)(nil).AuthorizeSecurityGroupIngressWithContext), arg0, arg1)
}

// CreateSecurityGroupWithContext mocks base method.
func (m *MockEC2) CreateSecurityGroupWithContext(arg0 context.Context, arg1 *ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecurityGroupWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.CreateSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityGroupWithContext indicates an expected call of CreateSecurityGroupWithContext.
func (mr *MockEC2MockRecorder) CreateSecurityGroupWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityGroupWithContext", reflect.TypeOf((*MockEC2)(nil).CreateSecurityGroupWithContext), arg0, arg1)
}

// CreateTagsWithContext mocks base method.
func (m *MockEC2) CreateTagsWithContext(arg0 context.Context, arg1 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTagsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTagsWithContext indicates an expected call of CreateTagsWithContext.
func (mr *MockEC2MockRecorder) CreateTagsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTagsWithContext", reflect.TypeOf((*MockEC2)(nil).CreateTagsWithContext), arg0, arg1)
}

// DeleteSecurityGroupWithContext mocks base method.
func (m *MockEC2) DeleteSecurityGroupWithContext(arg0 context.Context, arg1 *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityGroupWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.DeleteSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecurityGroupWithContext indicates an expected call of DeleteSecurityGroupWithContext.
func (mr *MockEC2MockRecorder) DeleteSecurityGroupWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityGroupWithContext", reflect.TypeOf((*MockEC2)(nil).DeleteSecurityGroupWithContext), arg0, arg1)
}

// DeleteTagsWithContext mocks base method.
func (m *MockEC2) DeleteTagsWithContext(arg0 context.Context, arg1 *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTagsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.DeleteTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTagsWithContext indicates an expected call of DeleteTagsWithContext.
func (mr *MockEC2MockRecorder) DeleteTagsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTagsWithContext", reflect.TypeOf((*MockEC2)(nil).DeleteTagsWithContext), arg0, arg1)
}

// DescribeAvailabilityZonesWithContext mocks base method.
func (m *MockEC2) DescribeAvailabilityZonesWithContext(arg0 context.Context, arg1 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAvailabilityZonesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.DescribeAvailabilityZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAvailabilityZonesWithContext indicates an expected call of DescribeAvailabilityZonesWithContext.
func (mr *MockEC2MockRecorder) DescribeAvailabilityZonesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailabilityZonesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeAvailabilityZonesWithContext), arg0, arg1)
}

// DescribeInstancesAsList mocks base method.
func (m *MockEC2) DescribeInstancesAsList(arg0 context.Context, arg1 *ec2.DescribeInstancesInput) ([]types.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstancesAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstancesAsList indicates an expected call of DescribeInstancesAsList.
func (mr *MockEC2MockRecorder) DescribeInstancesAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstancesAsList", reflect.TypeOf((*MockEC2)(nil).DescribeInstancesAsList), arg0, arg1)
}

// DescribeInstancesWithContext mocks base method.
func (m *MockEC2) DescribeInstancesWithContext(arg0 context.Context, arg1 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstancesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstancesWithContext indicates an expected call of DescribeInstancesWithContext.
func (mr *MockEC2MockRecorder) DescribeInstancesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstancesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeInstancesWithContext), arg0, arg1)
}

// DescribeNetworkInterfacesAsList mocks base method.
func (m *MockEC2) DescribeNetworkInterfacesAsList(arg0 context.Context, arg1 *ec2.DescribeNetworkInterfacesInput) ([]types.NetworkInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeNetworkInterfacesAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.NetworkInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNetworkInterfacesAsList indicates an expected call of DescribeNetworkInterfacesAsList.
func (mr *MockEC2MockRecorder) DescribeNetworkInterfacesAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNetworkInterfacesAsList", reflect.TypeOf((*MockEC2)(nil).DescribeNetworkInterfacesAsList), arg0, arg1)
}

// DescribeRouteTablesAsList mocks base method.
func (m *MockEC2) DescribeRouteTablesAsList(arg0 context.Context, arg1 *ec2.DescribeRouteTablesInput) ([]types.RouteTable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRouteTablesAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.RouteTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRouteTablesAsList indicates an expected call of DescribeRouteTablesAsList.
func (mr *MockEC2MockRecorder) DescribeRouteTablesAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRouteTablesAsList", reflect.TypeOf((*MockEC2)(nil).DescribeRouteTablesAsList), arg0, arg1)
}

// DescribeSecurityGroupsAsList mocks base method.
func (m *MockEC2) DescribeSecurityGroupsAsList(arg0 context.Context, arg1 *ec2.DescribeSecurityGroupsInput) ([]types.SecurityGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSecurityGroupsAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.SecurityGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityGroupsAsList indicates an expected call of DescribeSecurityGroupsAsList.
func (mr *MockEC2MockRecorder) DescribeSecurityGroupsAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroupsAsList", reflect.TypeOf((*MockEC2)(nil).DescribeSecurityGroupsAsList), arg0, arg1)
}

// DescribeSubnetsAsList mocks base method.
func (m *MockEC2) DescribeSubnetsAsList(arg0 context.Context, arg1 *ec2.DescribeSubnetsInput) ([]types.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSubnetsAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnetsAsList indicates an expected call of DescribeSubnetsAsList.
func (mr *MockEC2MockRecorder) DescribeSubnetsAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnetsAsList", reflect.TypeOf((*MockEC2)(nil).DescribeSubnetsAsList), arg0, arg1)
}

// DescribeVPCsAsList mocks base method.
func (m *MockEC2) DescribeVPCsAsList(arg0 context.Context, arg1 *ec2.DescribeVpcsInput) ([]types.Vpc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVPCsAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.Vpc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVPCsAsList indicates an expected call of DescribeVPCsAsList.
func (mr *MockEC2MockRecorder) DescribeVPCsAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVPCsAsList", reflect.TypeOf((*MockEC2)(nil).DescribeVPCsAsList), arg0, arg1)
}

// DescribeVpcsWithContext mocks base method.
func (m *MockEC2) DescribeVpcsWithContext(arg0 context.Context, arg1 *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.DescribeVpcsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcsWithContext indicates an expected call of DescribeVpcsWithContext.
func (mr *MockEC2MockRecorder) DescribeVpcsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcsWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeVpcsWithContext), arg0, arg1)
}

// RevokeSecurityGroupIngressWithContext mocks base method.
func (m *MockEC2) RevokeSecurityGroupIngressWithContext(arg0 context.Context, arg1 *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecurityGroupIngressWithContext", arg0, arg1)
	ret0, _ := ret[0].(*ec2.RevokeSecurityGroupIngressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeSecurityGroupIngressWithContext indicates an expected call of RevokeSecurityGroupIngressWithContext.
func (mr *MockEC2MockRecorder) RevokeSecurityGroupIngressWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecurityGroupIngressWithContext", reflect.TypeOf((*MockEC2)(nil).RevokeSecurityGroupIngressWithContext), arg0, arg1)
}

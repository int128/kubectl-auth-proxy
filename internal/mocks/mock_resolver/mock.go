// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/kauthproxy/internal/resolver (interfaces: FactoryInterface,Interface)
//
// Generated by this command:
//
//	mockgen -destination internal/mocks/mock_resolver/mock.go github.com/int128/kauthproxy/internal/resolver FactoryInterface,Interface
//

// Package mock_resolver is a generated GoMock package.
package mock_resolver

import (
	context "context"
	reflect "reflect"

	resolver "github.com/int128/kauthproxy/internal/resolver"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	rest "k8s.io/client-go/rest"
)

// MockFactoryInterface is a mock of FactoryInterface interface.
type MockFactoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryInterfaceMockRecorder
	isgomock struct{}
}

// MockFactoryInterfaceMockRecorder is the mock recorder for MockFactoryInterface.
type MockFactoryInterfaceMockRecorder struct {
	mock *MockFactoryInterface
}

// NewMockFactoryInterface creates a new mock instance.
func NewMockFactoryInterface(ctrl *gomock.Controller) *MockFactoryInterface {
	mock := &MockFactoryInterface{ctrl: ctrl}
	mock.recorder = &MockFactoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactoryInterface) EXPECT() *MockFactoryInterfaceMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockFactoryInterface) New(config *rest.Config) (resolver.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", config)
	ret0, _ := ret[0].(resolver.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockFactoryInterfaceMockRecorder) New(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockFactoryInterface)(nil).New), config)
}

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// FindPodByName mocks base method.
func (m *MockInterface) FindPodByName(ctx context.Context, namespace, podName string) (*v1.Pod, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPodByName", ctx, namespace, podName)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindPodByName indicates an expected call of FindPodByName.
func (mr *MockInterfaceMockRecorder) FindPodByName(ctx, namespace, podName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPodByName", reflect.TypeOf((*MockInterface)(nil).FindPodByName), ctx, namespace, podName)
}

// FindPodByServiceName mocks base method.
func (m *MockInterface) FindPodByServiceName(ctx context.Context, namespace, serviceName string) (*v1.Pod, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPodByServiceName", ctx, namespace, serviceName)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindPodByServiceName indicates an expected call of FindPodByServiceName.
func (mr *MockInterfaceMockRecorder) FindPodByServiceName(ctx, namespace, serviceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPodByServiceName", reflect.TypeOf((*MockInterface)(nil).FindPodByServiceName), ctx, namespace, serviceName)
}

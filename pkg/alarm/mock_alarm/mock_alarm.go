// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/thspinto/isecnet-go/pkg/alarm (interfaces: AlarmClient)

// Package mock_alarm is a generated GoMock package.
package mock_alarm

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	alarm "github.com/thspinto/isecnet-go/pkg/alarm"
)

// MockAlarmClient is a mock of AlarmClient interface.
type MockAlarmClient struct {
	ctrl     *gomock.Controller
	recorder *MockAlarmClientMockRecorder
}

// MockAlarmClientMockRecorder is the mock recorder for MockAlarmClient.
type MockAlarmClientMockRecorder struct {
	mock *MockAlarmClient
}

// NewMockAlarmClient creates a new mock instance.
func NewMockAlarmClient(ctrl *gomock.Controller) *MockAlarmClient {
	mock := &MockAlarmClient{ctrl: ctrl}
	mock.recorder = &MockAlarmClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlarmClient) EXPECT() *MockAlarmClientMockRecorder {
	return m.recorder
}

// GetPartialStatus mocks base method.
func (m *MockAlarmClient) GetPartialStatus(arg0 context.Context) (*alarm.StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartialStatus", arg0)
	ret0, _ := ret[0].(*alarm.StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartialStatus indicates an expected call of GetPartialStatus.
func (mr *MockAlarmClientMockRecorder) GetPartialStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartialStatus", reflect.TypeOf((*MockAlarmClient)(nil).GetPartialStatus), arg0)
}

// GetZones mocks base method.
func (m *MockAlarmClient) GetZones(arg0 context.Context, arg1 bool) ([]alarm.ZoneModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZones", arg0, arg1)
	ret0, _ := ret[0].([]alarm.ZoneModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZones indicates an expected call of GetZones.
func (mr *MockAlarmClientMockRecorder) GetZones(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZones", reflect.TypeOf((*MockAlarmClient)(nil).GetZones), arg0, arg1)
}

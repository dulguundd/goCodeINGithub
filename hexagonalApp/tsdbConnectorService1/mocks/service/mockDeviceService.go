// Code generated by MockGen. DO NOT EDIT.
// Source: deviceService.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"
	dto "tsdbConnectorService1/dto"

	errs "github.com/dulguundd/logError-lib/errs"
	gomock "github.com/golang/mock/gomock"
)

// MockDeviceService is a mock of DeviceService interface.
type MockDeviceService struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceMockRecorder
}

// MockDeviceServiceMockRecorder is the mock recorder for MockDeviceService.
type MockDeviceServiceMockRecorder struct {
	mock *MockDeviceService
}

// NewMockDeviceService creates a new mock instance.
func NewMockDeviceService(ctrl *gomock.Controller) *MockDeviceService {
	mock := &MockDeviceService{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceService) EXPECT() *MockDeviceServiceMockRecorder {
	return m.recorder
}

// GetAllDevice mocks base method.
func (m *MockDeviceService) GetAllDevice() ([]dto.DeviceResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDevice")
	ret0, _ := ret[0].([]dto.DeviceResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAllDevice indicates an expected call of GetAllDevice.
func (mr *MockDeviceServiceMockRecorder) GetAllDevice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDevice", reflect.TypeOf((*MockDeviceService)(nil).GetAllDevice))
}

// GetDevice mocks base method.
func (m *MockDeviceService) GetDevice(arg0 int) (*dto.DeviceResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevice", arg0)
	ret0, _ := ret[0].(*dto.DeviceResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice.
func (mr *MockDeviceServiceMockRecorder) GetDevice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockDeviceService)(nil).GetDevice), arg0)
}

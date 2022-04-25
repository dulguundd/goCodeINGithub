// Code generated by MockGen. DO NOT EDIT.
// Source: dataService.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"
	dto "tsdbConnectorService1/dto"

	errs "github.com/dulguundd/logError-lib/errs"
	gomock "github.com/golang/mock/gomock"
)

// MockDataService is a mock of DataService interface.
type MockDataService struct {
	ctrl     *gomock.Controller
	recorder *MockDataServiceMockRecorder
}

// MockDataServiceMockRecorder is the mock recorder for MockDataService.
type MockDataServiceMockRecorder struct {
	mock *MockDataService
}

// NewMockDataService creates a new mock instance.
func NewMockDataService(ctrl *gomock.Controller) *MockDataService {
	mock := &MockDataService{ctrl: ctrl}
	mock.recorder = &MockDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataService) EXPECT() *MockDataServiceMockRecorder {
	return m.recorder
}

// GetLastDataOfTodayByHour mocks base method.
func (m *MockDataService) GetLastDataOfTodayByHour() ([]dto.DataResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastDataOfTodayByHour")
	ret0, _ := ret[0].([]dto.DataResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetLastDataOfTodayByHour indicates an expected call of GetLastDataOfTodayByHour.
func (mr *MockDataServiceMockRecorder) GetLastDataOfTodayByHour() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastDataOfTodayByHour", reflect.TypeOf((*MockDataService)(nil).GetLastDataOfTodayByHour))
}

// GetLastDataOfTodayByHourOfDevice mocks base method.
func (m *MockDataService) GetLastDataOfTodayByHourOfDevice(id int) ([]dto.DataResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastDataOfTodayByHourOfDevice", id)
	ret0, _ := ret[0].([]dto.DataResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetLastDataOfTodayByHourOfDevice indicates an expected call of GetLastDataOfTodayByHourOfDevice.
func (mr *MockDataServiceMockRecorder) GetLastDataOfTodayByHourOfDevice(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastDataOfTodayByHourOfDevice", reflect.TypeOf((*MockDataService)(nil).GetLastDataOfTodayByHourOfDevice), id)
}

// PostData mocks base method.
func (m *MockDataService) PostData(req dto.NewDataRequest) (*dto.DataResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostData", req)
	ret0, _ := ret[0].(*dto.DataResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// PostData indicates an expected call of PostData.
func (mr *MockDataServiceMockRecorder) PostData(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostData", reflect.TypeOf((*MockDataService)(nil).PostData), req)
}
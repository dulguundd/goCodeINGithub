// Code generated by MockGen. DO NOT EDIT.
// Source: data.go

// Package data is a generated GoMock package.
package data

import (
	reflect "reflect"
	data "tsdbConnectorService1/domain/data"

	errs "github.com/dulguundd/logError-lib/errs"
	gomock "github.com/golang/mock/gomock"
)

// MockDataRepository is a mock of DataRepository interface.
type MockDataRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDataRepositoryMockRecorder
}

// MockDataRepositoryMockRecorder is the mock recorder for MockDataRepository.
type MockDataRepositoryMockRecorder struct {
	mock *MockDataRepository
}

// NewMockDataRepository creates a new mock instance.
func NewMockDataRepository(ctrl *gomock.Controller) *MockDataRepository {
	mock := &MockDataRepository{ctrl: ctrl}
	mock.recorder = &MockDataRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataRepository) EXPECT() *MockDataRepositoryMockRecorder {
	return m.recorder
}

// GetLastDataOfTodayByHour mocks base method.
func (m *MockDataRepository) GetLastDataOfTodayByHour() ([]data.Data, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastDataOfTodayByHour")
	ret0, _ := ret[0].([]data.Data)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetLastDataOfTodayByHour indicates an expected call of GetLastDataOfTodayByHour.
func (mr *MockDataRepositoryMockRecorder) GetLastDataOfTodayByHour() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastDataOfTodayByHour", reflect.TypeOf((*MockDataRepository)(nil).GetLastDataOfTodayByHour))
}

// GetLastDataOfTodayByHourOfDevice mocks base method.
func (m *MockDataRepository) GetLastDataOfTodayByHourOfDevice(id int) ([]data.Data, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastDataOfTodayByHourOfDevice", id)
	ret0, _ := ret[0].([]data.Data)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetLastDataOfTodayByHourOfDevice indicates an expected call of GetLastDataOfTodayByHourOfDevice.
func (mr *MockDataRepositoryMockRecorder) GetLastDataOfTodayByHourOfDevice(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastDataOfTodayByHourOfDevice", reflect.TypeOf((*MockDataRepository)(nil).GetLastDataOfTodayByHourOfDevice), id)
}

// PostData mocks base method.
func (m *MockDataRepository) PostData(newData data.Data) (*data.Data, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostData", newData)
	ret0, _ := ret[0].(*data.Data)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// PostData indicates an expected call of PostData.
func (mr *MockDataRepositoryMockRecorder) PostData(newData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostData", reflect.TypeOf((*MockDataRepository)(nil).PostData), newData)
}

package app

import (
	"github.com/dulguundd/logError-lib/errs"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"tsdbConnectorService1/dto"
	"tsdbConnectorService1/mocks/service"
)

var deviceRouter *mux.Router
var deh DeviceHandlers
var mockDeviceService *service.MockDeviceService

func setupDevice(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockDeviceService = service.NewMockDeviceService(ctrl)
	deh = DeviceHandlers{mockDeviceService}
	deviceRouter = mux.NewRouter()
	deviceRouter.HandleFunc("/device", deh.getAllDevice)
	deviceRouter.HandleFunc("/device/{device_id:[0-9]+}", deh.getDevice)
	return func() {
		deviceRouter = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_devices_with_status_code_200_getAllDevice(t *testing.T) {
	//Arrange
	teardown := setupDevice(t)
	defer teardown()

	dummyDevices := []dto.DeviceResponse{
		{1, "device1", "device1", "0000001", "1"},
		{2, "device2", "device2", "0000002", "0"},
	}
	mockDeviceService.EXPECT().GetAllDevice().Return(dummyDevices, nil)

	request, _ := http.NewRequest(http.MethodGet, "/device", nil)

	//Act
	recorder := httptest.NewRecorder()
	deviceRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_getAllDevice(t *testing.T) {
	//Arrange
	teardown := setupDevice(t)
	defer teardown()

	mockDeviceService.EXPECT().GetAllDevice().Return(nil, errs.NewUnexpectedError("some database error"))

	request, _ := http.NewRequest(http.MethodGet, "/device", nil)

	//Act
	recoder := httptest.NewRecorder()
	deviceRouter.ServeHTTP(recoder, request)

	//Assert
	if recoder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_devices_with_status_code_200_getDevice(t *testing.T) {
	//Arrange
	teardown := setupDevice(t)
	defer teardown()

	dummyDevices := dto.DeviceResponse{1, "device1", "device1", "0000001", "1"}

	deviceid := 1
	mockDeviceService.EXPECT().GetDevice(deviceid).Return(&dummyDevices, nil)

	request, _ := http.NewRequest(http.MethodGet, "/device/1", nil)

	//Act
	recorder := httptest.NewRecorder()
	deviceRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_getDevice(t *testing.T) {
	//Arrange
	teardown := setupDevice(t)
	defer teardown()

	deviceid := 1
	mockDeviceService.EXPECT().GetDevice(deviceid).Return(nil, errs.NewUnexpectedError("some database error"))

	request, _ := http.NewRequest(http.MethodGet, "/device/1", nil)

	//Act
	recoder := httptest.NewRecorder()
	deviceRouter.ServeHTTP(recoder, request)

	//Assert
	if recoder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

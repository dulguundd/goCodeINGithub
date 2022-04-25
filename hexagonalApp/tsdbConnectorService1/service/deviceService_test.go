package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"github.com/golang/mock/gomock"
	"testing"
	realdevice "tsdbConnectorService1/domain/device"
	"tsdbConnectorService1/mocks/domain/device"
)

var mockRepoDevice *device.MockDeviceRepository
var serviceDevice DeviceService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepoDevice = device.NewMockDeviceRepository(ctrl)
	serviceDevice = NewDeviceService(mockRepoDevice)
	return func() {
		serviceDevice = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_error_with_database_error_GetAllDevice(t *testing.T) {
	//Arrange
	teardown := setup(t)
	defer teardown()
	mockRepoDevice.EXPECT().GetAllDevice().Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	//Act
	_, appError := serviceDevice.GetAllDevice()

	//Assert
	if appError == nil {
		t.Error("Test failed didnt return database error")
	}
}

func Test_should_return_devices_GetAllDevice(t *testing.T) {
	//Arrange
	teardown := setup(t)
	defer teardown()
	dummyDevice := []realdevice.Device{
		{3, "device3", "device3", "100217", 1},
		{2, "device2", "device2", "100216", 1},
		{1, "device1", "device1", "100215", 0},
	}
	mockRepoDevice.EXPECT().GetAllDevice().Return(dummyDevice, nil)

	//Act
	res, appError := serviceDevice.GetAllDevice()

	//Assert
	if appError != nil {
		t.Error("Test failed error is not null")
	}
	if res == nil {
		t.Error("Test failed no device return")
	}
}

func Test_should_return_error_with_database_error_GetDevice(t *testing.T) {
	//Arrange
	teardown := setup(t)
	defer teardown()
	deviceid := 1
	mockRepoDevice.EXPECT().GetDevice(deviceid).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	//Act
	_, appError := serviceDevice.GetDevice(deviceid)

	//Assert
	if appError == nil {
		t.Error("Test failed didnt return database error")
	}
}

func Test_should_return_device_GetDevice(t *testing.T) {
	//Arrange
	teardown := setup(t)
	defer teardown()
	deviceid := 1
	dummyDevice := realdevice.Device{Device_id: 1, Device_name: "device1", Device_spec: "device1", Serial_id: "100217", Status: 1}

	mockRepoDevice.EXPECT().GetDevice(deviceid).Return(&dummyDevice, nil)

	//Act
	res, appError := serviceDevice.GetDevice(deviceid)

	//Assert
	if appError != nil {
		t.Error("Test failed error is not null")
	}
	if res == nil {
		t.Error("Test failed no device return")
	}
	if res.DeviceId != deviceid {
		t.Error("Test failed device from database with another device data")
	}
}

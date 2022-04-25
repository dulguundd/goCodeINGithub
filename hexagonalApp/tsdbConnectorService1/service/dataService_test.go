package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
	realdata "tsdbConnectorService1/domain/data"
	"tsdbConnectorService1/dto"
	"tsdbConnectorService1/mocks/domain/data"
)

func Test_should_return_a_validation_error_response_when_request_is_not_validated_PostData(t *testing.T) {
	//Arrange
	request := dto.NewDataRequest{
		Temp:      150,
		Device_Id: 3,
	}
	service := NewDataService(nil)

	//Act
	_, appError := service.PostData(request)

	//Arssert
	if appError == nil {
		t.Error("Failed while testing the new data validation")
	}
}

var mockRepoData *data.MockDataRepository
var serviceData DataService

func setupData(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepoData = data.NewMockDataRepository(ctrl)
	serviceData = NewDataService(mockRepoData)
	return func() {
		serviceData = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannot_be_created_PostData(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()

	req := dto.NewDataRequest{
		Temp:      50,
		Device_Id: 3,
	}
	newData := realdata.Data{
		Time:      time.Now(),
		Temp:      req.Temp,
		Device_id: req.Device_Id,
	}

	mockRepoData.EXPECT().PostData(newData).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	//Act
	_, appError := serviceData.PostData(req)

	//Assert
	if appError == nil {
		t.Error("Test failed while validating error for new data")
	}
}

func Test_should_return_new_data_response_when_a_new_data_is_saved_successfully_PostData(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()
	req := dto.NewDataRequest{
		Temp:      50,
		Device_Id: 3,
	}
	newData := realdata.Data{
		Time:      time.Now(),
		Temp:      req.Temp,
		Device_id: req.Device_Id,
	}
	mockRepoData.EXPECT().PostData(newData).Return(&newData, nil)

	//Act
	savedData, appError := serviceData.PostData(req)

	//Assert
	if appError != nil {
		t.Error("Test failed while saving new data")
	}
	if savedData.Temp != newData.Temp {
		t.Error("Test failed while new data didnt match saved data")
	}
}

func Test_should_return_error_with_database_error_GetLastDataOfTodayByHour(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()

	mockRepoData.EXPECT().GetLastDataOfTodayByHour().Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	//Act
	_, appError := serviceData.GetLastDataOfTodayByHour()

	//Assert
	if appError == nil {
		t.Error("Test failed didnt return database error")
	}
}

func Test_should_return_data_GetLastDataOfTodayByHour(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()

	dummyData := []realdata.Data{
		{time.Now(), 3, 2},
		{time.Now(), 16, 1},
		{time.Now(), 10, 2},
	}
	mockRepoData.EXPECT().GetLastDataOfTodayByHour().Return(dummyData, nil)

	//Act
	data, appError := serviceData.GetLastDataOfTodayByHour()

	//Assert
	if appError != nil {
		t.Error("Test failed error is not null")
	}
	if data == nil {
		t.Error("Test failed no data return")
	}
}

func Test_should_return_error_with_database_error_GetLastDataOfTodayByHourOfDevice(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()
	var deviceid int = 1
	mockRepoData.EXPECT().GetLastDataOfTodayByHourOfDevice(deviceid).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	//Act
	_, appError := serviceData.GetLastDataOfTodayByHourOfDevice(deviceid)

	//Assert
	if appError == nil {
		t.Error("Test failed didnt return database error")
	}
}

func Test_should_return_data_GetLastDataOfTodayByHourOfDevice(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()
	var deviceid int = 1
	dummyData := []realdata.Data{
		{time.Now(), 3, 1},
		{time.Now(), 16, 1},
		{time.Now(), 10, 1},
	}
	mockRepoData.EXPECT().GetLastDataOfTodayByHourOfDevice(deviceid).Return(dummyData, nil)

	var data []dto.DataResponse
	var appError *errs.AppError

	//Act
	data, appError = serviceData.GetLastDataOfTodayByHourOfDevice(deviceid)

	//Assert
	if appError != nil {
		t.Error("Test failed error is not null")
	}
	if data == nil {
		t.Error("Test failed no data return")
	}
	for i := range data {
		if data[i].Device_Id != deviceid {
			t.Error("Test failed data from database with another device data")
		}
	}
}

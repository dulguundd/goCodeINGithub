package app

import (
	"bytes"
	"encoding/json"
	"github.com/dulguundd/logError-lib/errs"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"tsdbConnectorService1/dto"
	"tsdbConnectorService1/mocks/service"
)

var dataRouter *mux.Router
var dah DataHandlers
var mockDataService *service.MockDataService

func setupData(t *testing.T) func() {
	datactrl := gomock.NewController(t)
	mockDataService = service.NewMockDataService(datactrl)
	dah = DataHandlers{mockDataService}
	dataRouter = mux.NewRouter()
	dataRouter.HandleFunc("/data/today", dah.GetLastDataOfTodayByHour).Methods(http.MethodGet)
	dataRouter.HandleFunc("/data/today/{device_id:[0-9]+}", dah.GetLastDataOfTodayByHourOfDevice).Methods(http.MethodGet)
	dataRouter.HandleFunc("/data/{device_id:[0-9]+}/{temp:[0-9]+}", dah.PostData).Methods(http.MethodPost)
	dataRouter.HandleFunc("/data", dah.PostDataBody).Methods(http.MethodPost)
	return func() {
		dataRouter = nil
		defer datactrl.Finish()
	}
}

func Test_should_return_data_with_status_code_200_GetLastDataOfTodayByHour(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()

	dummyData := []dto.DataResponse{
		{time.Now(), 10, 1},
		{time.Now(), 20, 2},
	}
	mockDataService.EXPECT().GetLastDataOfTodayByHour().Return(dummyData, nil)

	request, appError := http.NewRequest(http.MethodGet, "/data/today", nil)

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
	if appError != nil {
		t.Error("Failed while testing the error should null")
	}
}

func Test_should_return_status_code_500_with_error_message_GetLastDataOfTodayByHour(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()

	mockDataService.EXPECT().GetLastDataOfTodayByHour().Return(nil, errs.NewUnexpectedError("some database error"))

	request, _ := http.NewRequest(http.MethodGet, "/data/today", nil)

	//Act
	recoder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recoder, request)

	//Assert
	if recoder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
	if recoder.Body == nil {
		t.Error("Failed, Response body is null")
	}
}

func Test_should_return_data_with_status_code_200_GetLastDataOfTodayByHourOfDevice(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()

	dummyData := []dto.DataResponse{
		{time.Now(), 10, 1},
		{time.Now(), 20, 2},
	}
	deviceid := 1
	mockDataService.EXPECT().GetLastDataOfTodayByHourOfDevice(deviceid).Return(dummyData, nil)
	request, appError := http.NewRequest(http.MethodGet, "/data/today/1", nil)

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
	if appError != nil {
		t.Error("Failed while testing the error should null")
	}
}

func Test_should_return_status_code_500_with_error_message_GetLastDataOfTodayByHourOfDevice(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()
	deviceid := 1
	mockDataService.EXPECT().GetLastDataOfTodayByHourOfDevice(deviceid).Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/data/today/1", nil)

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
	if recorder.Body == nil {
		t.Error("Failed, Response body is null")
	}
}

func Test_should_return_data_with_status_code_200_PostData(t *testing.T) {
	teardown := setupData(t)
	defer teardown()
	req := dto.NewDataRequest{
		Temp:      10,
		Device_Id: 1,
	}
	dummyData := dto.DataResponse{
		Time:      time.Now(),
		Temp:      req.Temp,
		Device_Id: req.Device_Id,
	}
	mockDataService.EXPECT().PostData(req).Return(&dummyData, nil)
	request, appError := http.NewRequest(http.MethodPost, "/data/1/10", nil)

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
	if appError != nil {
		t.Error("Failed while testing the error should null")
	}
}

func Test_should_return_status_code_500_with_error_message_PostData(t *testing.T) {
	//Arrange
	teardown := setupData(t)
	defer teardown()
	req := dto.NewDataRequest{
		Temp:      10,
		Device_Id: 1,
	}
	mockDataService.EXPECT().PostData(req).Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodPost, "/data/1/10", nil)

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
	if recorder.Body == nil {
		t.Error("Failed, Response body is null")
	}
}

func Test_should_return_data_with_status_code_200_PostDataBody(t *testing.T) {
	teardown := setupData(t)
	defer teardown()
	req := dto.NewDataRequest{
		Temp:      10,
		Device_Id: 1,
	}
	dummyData := dto.DataResponse{
		Time:      time.Now(),
		Temp:      req.Temp,
		Device_Id: req.Device_Id,
	}
	mockDataService.EXPECT().PostData(req).Return(&dummyData, nil)
	requestBody, _ := json.Marshal(map[string]int{
		"temp":      10,
		"device_id": 1,
	})
	request, appError := http.NewRequest(http.MethodPost, "/data",
		bytes.NewBuffer(requestBody))

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
	if appError != nil {
		t.Error("Failed while testing the error should null")
	}
}

func Test_should_return_status_code_500_with_error_message_PostDataBody(t *testing.T) {
	teardown := setupData(t)
	defer teardown()
	req := dto.NewDataRequest{
		Temp:      10,
		Device_Id: 1,
	}
	mockDataService.EXPECT().PostData(req).Return(nil, errs.NewUnexpectedError("some database error"))
	requestBody, _ := json.Marshal(map[string]int{
		"temp":      10,
		"device_id": 1,
	})
	request, _ := http.NewRequest(http.MethodPost, "/data",
		bytes.NewBuffer(requestBody))

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
	if recorder.Body == nil {
		t.Error("Failed, Response body is null")
	}
}

func Test_should_return_status_code_400_with_error_message_PostDataBody(t *testing.T) {
	teardown := setupData(t)
	defer teardown()
	type requestBody struct {
		Temp     int    `json:"temp"`
		DeviceId string `json:"device_id"`
	}
	rawReqBody := &requestBody{
		Temp:     10,
		DeviceId: "1",
	}
	reqBody, _ := json.Marshal(rawReqBody)
	request, _ := http.NewRequest(http.MethodPost, "/data",
		bytes.NewBuffer(reqBody))

	//Act
	recorder := httptest.NewRecorder()
	dataRouter.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusBadRequest {
		t.Error("Failed while testing the status code")
	}
	if recorder.Body == nil {
		t.Error("Failed, Response body is null")
	}
}

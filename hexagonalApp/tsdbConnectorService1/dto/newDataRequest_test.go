package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_new_data_temperature_is_not_correct(t *testing.T) {
	//Arrange
	request := NewDataRequest{
		Temp: -100,
	}

	//Act
	appError := request.Validate()

	//Assert
	if appError.Message != "Check temperature value" {
		t.Error("Invalid message while testing New Data Request")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing New Data Request")
	}
}

func Test_should_return_response_when_new_data_temperature_is_correct(t *testing.T) {
	//Arrange
	request := NewDataRequest{
		Temp: 10,
	}

	//Act

	appError := request.Validate()

	if appError != nil {
		t.Error("Error while reading correct data")
	}
}

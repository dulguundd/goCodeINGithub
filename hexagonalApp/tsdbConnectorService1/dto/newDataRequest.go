package dto

import (
	"github.com/dulguundd/logError-lib/errs"
)

type NewDataRequest struct {
	Temp      float64 `json:"temp"`
	Device_Id int     `json:"device_id"`
}

func (r NewDataRequest) Validate() *errs.AppError {
	if r.Temp > -100 && r.Temp < 100 {
		return nil
	} else {
		return errs.NewValidationError("Check temperature value")
	}
}

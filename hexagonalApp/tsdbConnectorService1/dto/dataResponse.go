package dto

import "time"

type DataResponse struct {
	Time      time.Time `json:"time"`
	Temp      float64   `json:"temp"`
	Device_Id int       `json:"device_id"`
}

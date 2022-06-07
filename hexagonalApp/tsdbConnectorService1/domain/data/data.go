package data

import (
	"github.com/dulguundd/logError-lib/errs"
	"time"
	"tsdbConnectorService1/dto"
)

type Data struct {
	Bucket    time.Time `json:"bucket"`
	Temp      float64   `json:"temp"`
	Device_id int       `json:"device_id"`
}

//go:generate mockgen -destination=../../mocks/domain/data/mockDataRepository.go -package=data -source=data.go DataRepository

type DataRepository interface {
	GetLastDataOfTodayByHour() ([]Data, *errs.AppError)
	GetLastDataOfTodayByHourOfDevice(id int) ([]Data, *errs.AppError)
	PostData(newData Data) (*Data, *errs.AppError)
}

func (d Data) ToDtoData() dto.DataResponse {
	return dto.DataResponse{
		Time:      d.Bucket,
		Temp:      d.Temp,
		Device_Id: d.Device_id,
	}
}

func ReqToNewData(Temp float64, Device_Id int) Data {
	return Data{
		Bucket:    time.Now(),
		Temp:      Temp,
		Device_id: Device_Id,
	}
}

func (d Data) PostedDataToDtoData() *dto.DataResponse {
	return &dto.DataResponse{
		Time:      d.Bucket,
		Temp:      d.Temp,
		Device_Id: d.Device_id,
	}
}

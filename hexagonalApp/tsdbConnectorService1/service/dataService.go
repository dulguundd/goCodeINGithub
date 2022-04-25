package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"tsdbConnectorService1/domain/data"
	"tsdbConnectorService1/dto"
)

//go:generate mockgen -destination=../mocks/service/mockDataService.go -package=service -source=dataService.go Dataservice

type DataService interface {
	GetLastDataOfTodayByHour() ([]dto.DataResponse, *errs.AppError)
	GetLastDataOfTodayByHourOfDevice(id int) ([]dto.DataResponse, *errs.AppError)
	PostData(req dto.NewDataRequest) (*dto.DataResponse, *errs.AppError)
}

type DefaultDataService struct {
	repo data.DataRepository
}

func (s DefaultDataService) GetLastDataOfTodayByHour() ([]dto.DataResponse, *errs.AppError) {
	if d, err := s.repo.GetLastDataOfTodayByHour(); err != nil {
		return nil, err
	} else {
		var response []dto.DataResponse
		for i := range d {
			response = append(response, d[i].ToDtoData())
		}
		return response, nil
	}
}

func (s DefaultDataService) GetLastDataOfTodayByHourOfDevice(id int) ([]dto.DataResponse, *errs.AppError) {
	if d, err := s.repo.GetLastDataOfTodayByHourOfDevice(id); err != nil {
		return nil, err
	} else {
		var response []dto.DataResponse
		for i := range d {
			response = append(response, d[i].ToDtoData())
		}
		return response, nil
	}
}

func (s DefaultDataService) PostData(req dto.NewDataRequest) (*dto.DataResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	newData := data.ReqToNewData(req.Temp, req.Device_Id)
	if savedData, err := s.repo.PostData(newData); err != nil {
		return nil, err
	} else {
		return savedData.PostedDataToDtoData(), nil
	}
}

func NewDataService(repository data.DataRepository) DefaultDataService {
	return DefaultDataService{(repository)}
}

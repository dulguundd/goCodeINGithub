package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"tsdbConnectorService1/domain/device"
	"tsdbConnectorService1/dto"
)

//go:generate mockgen -destination=../mocks/service/mockDeviceService.go -package=service -source=deviceService.go Deviceservice

type DeviceService interface {
	GetAllDevice() ([]dto.DeviceResponse, *errs.AppError)
	GetDevice(int) (*dto.DeviceResponse, *errs.AppError)
}

type DefaultDeviceService struct {
	repo device.DeviceRepository
}

func (s DefaultDeviceService) GetAllDevice() ([]dto.DeviceResponse, *errs.AppError) {
	d, err := s.repo.GetAllDevice()
	if err != nil {
		return nil, err
	}
	var response []dto.DeviceResponse
	for i := range d {
		response = append(response, d[i].ToDtoDevice())
	}
	return response, nil
}

func (s DefaultDeviceService) GetDevice(id int) (*dto.DeviceResponse, *errs.AppError) {
	d, err := s.repo.GetDevice(id)
	if err != nil {
		return nil, err
	}
	response := d.ToDtoDevice()
	return &response, nil
}

func NewDeviceService(repository device.DeviceRepository) DefaultDeviceService {
	return DefaultDeviceService{(repository)}
}

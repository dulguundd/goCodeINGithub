package device

import (
	"github.com/dulguundd/logError-lib/errs"
	"tsdbConnectorService1/dto"
)

//go:generate mockgen -destination=../../mocks/domain/device/mockDeviceRepository.go -package=device -source=device.go DeviceRepository

type Device struct {
	Device_id   int    `json:"device_id"`
	Device_name string `json:"device_name"`
	Device_spec string `json:"device_spec"`
	Serial_id   string `json:"serial_id"`
	Status      int    `json:"status"`
}

func (d Device) statusAsText() string {
	statusAsText := "active"
	if d.Status == 0 {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (d Device) ToDtoDevice() dto.DeviceResponse {
	return dto.DeviceResponse{
		DeviceId:   d.Device_id,
		DeviceName: d.Device_name,
		DeviceSpec: d.Device_spec,
		SerialId:   d.Serial_id,
		Status:     d.statusAsText(),
	}
}

type DeviceRepository interface {
	GetAllDevice() ([]Device, *errs.AppError)
	GetDevice(id int) (*Device, *errs.AppError)
}

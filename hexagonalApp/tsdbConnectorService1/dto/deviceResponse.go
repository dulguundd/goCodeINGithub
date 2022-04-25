package dto

type DeviceResponse struct {
	DeviceId   int    `json:"device_id"`
	DeviceName string `json:"device_name"`
	DeviceSpec string `json:"device_spec"`
	SerialId   string `json:"serial_id"`
	Status     string `json:"status"`
}

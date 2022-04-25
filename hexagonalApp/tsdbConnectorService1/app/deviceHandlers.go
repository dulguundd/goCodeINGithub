package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
	"tsdbConnectorService1/service"
)

type DeviceHandlers struct {
	service service.DeviceService
}

func (ch *DeviceHandlers) getAllDevice(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	devices, err := ch.service.GetAllDevice()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, devices)
		serviceLatencyLogger(start)
	}
}

func (ch *DeviceHandlers) getDevice(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	id := vars["device_id"]

	idInt, _ := strconv.Atoi(id)

	device, err := ch.service.GetDevice(idInt)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, device)
		serviceLatencyLogger(start)
	}
}

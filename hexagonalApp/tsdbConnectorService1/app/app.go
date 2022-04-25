package app

import (
	"encoding/json"
	"fmt"
	"github.com/dulguundd/logError-lib/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"tsdbConnectorService1/domain/data"
	"tsdbConnectorService1/domain/device"
	"tsdbConnectorService1/service"
)

func Start() {
	serviceConfig := sanityCheckService()

	router := mux.NewRouter()

	// wiring
	dbClient := getDbClient()
	dataRepositoryDb := data.NewRepositoryDb(dbClient)
	deviceRepositoryDb := device.NewRepositoryDb(dbClient)

	deh := DeviceHandlers{service.NewDeviceService(deviceRepositoryDb)}
	dah := DataHandlers{service.NewDataService(dataRepositoryDb)}

	//define device routes
	router.HandleFunc("/device", deh.getAllDevice).Methods(http.MethodGet)
	router.HandleFunc("/device/{device_id:[0-9]+}", deh.getDevice).Methods(http.MethodGet)

	//define data routes
	router.HandleFunc("/data/today", dah.GetLastDataOfTodayByHour).Methods(http.MethodGet)
	router.HandleFunc("/data/today/{device_id:[0-9]+}", dah.GetLastDataOfTodayByHourOfDevice).Methods(http.MethodGet)
	router.HandleFunc("/data/{device_id:[0-9]+}/{temp:[0-9]+}", dah.PostData).Methods(http.MethodPost)
	router.HandleFunc("/data", dah.PostDataBody).Methods(http.MethodPost)
	// connectDB()

	//starting server
	log.Fatal(http.ListenAndServe(serviceConfig.address+":"+serviceConfig.port, router))
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	logger.Info(logMessage)
}

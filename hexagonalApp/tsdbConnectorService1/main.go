package main

import (
	"github.com/dulguundd/logError-lib/logger"
	"tsdbConnectorService1/app"
)

func main() {
	logger.Info("Starting the application.....")
	app.Start()
	logger.Info("Running")
}

package main

import (
	"fmt"

	Loader "github.com/pi-prakhar/utils/loader"
	Log "github.com/pi-prakhar/utils/logger"
)

var (
	logger Log.Logger
)

func init() {
	serviceName, err := Loader.GetValueFromConf("service_name")

	logger = Log.New(Log.DEBUG, serviceName)
	if err != nil {
		logger.Warn(fmt.Sprintf("%s", err))
	}
}

func main() {

	logger.Debug("This is a debug message")
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")

}

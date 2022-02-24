package shared

import (
	"log"

	"go.uber.org/zap"
)

func ZapLogCustom(message string, errorlevel string) {

	configLogger := zap.NewProductionConfig()
	configLogger.OutputPaths = []string{"stdout"}
	ZapLogger, err := configLogger.Build()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	defer ZapLogger.Sync()

	if errorlevel == "info" {
		ZapLogger.Info(message)
	} else if errorlevel == "warn" {
		ZapLogger.Warn(message)
	} else if errorlevel == "error" {
		ZapLogger.Error(message)
	} else if errorlevel == "fatal" {
		ZapLogger.Fatal(message)
	}

}

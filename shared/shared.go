package shared

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapLogCustom(message []string, errorlevel string) {

	configLogger := zap.NewProductionConfig()
	configLogger.EncoderConfig.TimeKey = "ts"
	configLogger.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	configLogger.OutputPaths = []string{"stdout"}
	ZapLogger, err := configLogger.Build()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	defer ZapLogger.Sync()

	if errorlevel == "info" {
		ZapLogger.Info(message[0])
	} else if errorlevel == "warn" {
		ZapLogger.Warn(message[0])
	} else if errorlevel == "error" {
		ZapLogger.Error(message[0])
	} else if errorlevel == "fatal" {
		ZapLogger.Fatal(message[0])
	}

}

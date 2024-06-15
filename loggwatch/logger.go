package loggwatch

import (
	"watch/confwatch"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger() (*zap.Logger, error) {
	var errorFile, _ = confwatch.LoadConfig("../config.json")

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{errorFile.LogErrorFile} // Nome do arquivo de log
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		MessageKey:     "message",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	return config.Build()
}

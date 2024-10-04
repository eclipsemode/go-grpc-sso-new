package logger

import (
	"go.uber.org/zap"
)

func New() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	sugar := logger.Sugar()

	sugar.Info("Logger initialized")

	return sugar
}

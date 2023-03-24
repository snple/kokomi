package log

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init(debug bool) {
	if debug {
		Logger, _ = zap.NewDevelopment()
	} else {
		Logger, _ = zap.NewProduction()
	}
}

func Sync() {
	Logger.Sync()
}

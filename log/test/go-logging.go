package test

import (
	logger "github.com/op/go-logging"
)

func test() {
	logger.Error("test error")
	logger.Info("test info")
	logger.Debug("test debug")

}

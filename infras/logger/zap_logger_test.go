package logger_test

import (
	"encoding/json"
	"errors"
	"testing"

	"GoWeb/infras/configs"
	"GoWeb/infras/logger"
)

// test logger
func TestApiLogger_Logger(t *testing.T) {
	cfg := &configs.Config{
		Logger: configs.Logger{
			Development:   true,
			DisableCaller: false,
			Encoding:      "console",
			Level:         "info",
			EnableFile:    false,
			InfoFileName:  "./logs/server.log",
			MaxBackups:    5,
			MaxAge:        30,
			MaxSize:       1,
		},
	}
	copy := cfg
	data, _ := json.Marshal(copy)
	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()
	logger.Info("info level test")
	logger.Error("121212121212 error")
	logger.Errorf("checker fatal error: %v", errors.New(string(data)))
	logger.Warn("warn level test")
	logger.Debug("debug level test")
	logger.Infof("info level test: %s", "111")
	logger.Errorf("error level test: %s", "111")
	logger.Warnf("warn level test: %s", "111")
	logger.Debugf("debug level test: %s", "111")
}

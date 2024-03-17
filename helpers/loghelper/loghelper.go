package loghelper

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// Set log level
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel) // Default level
	}
	// Set log format
	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			// Customize the time format
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		// Default to text formatter
		logger.SetFormatter(&logrus.TextFormatter{
			// Customize the time format
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}
	return logger
}

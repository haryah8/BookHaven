package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger initializes and configures logrus
func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// Set log level
	logger.SetLevel(logrus.InfoLevel)

	// Set log format with timestamp and colors
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,                  // Force colors in output
		DisableTimestamp: false,                 // Timestamp included
		TimestampFormat:  "2006-01-02 15:04:05", // Custom timestamp format
		FullTimestamp:    true,
		// Optionally, you can set DisableQuote: true if you want to remove quotes around log fields
	})

	// Set output destination (e.g., stdout)
	logger.SetOutput(os.Stdout)

	return logger
}

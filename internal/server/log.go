package server

import (
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type ServerLogger struct {
	Logger *logrus.Logger
}

// Example: [IP] [POST] [/signup] ..
func initServerLogger() *ServerLogger {
	serverLogger := &ServerLogger{}
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&formatter.Formatter{
		FieldsOrder:     []string{"source", "method", "endpoint"},
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        false,
		ShowFullLevel:   true,
		HideKeys:        true,
	})
	serverLogger.Logger = logger
	return serverLogger
}

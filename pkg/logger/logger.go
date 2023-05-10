package logger

import (
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Driver *logrus.Logger
}

func NewLogger() *Logger {
	logger := &Logger{}
	driver := logrus.New()
	driver.SetOutput(os.Stdout)
	driver.SetLevel(logrus.DebugLevel)
	driver.SetFormatter(&formatter.Formatter{
		FieldsOrder:     []string{"source"},
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        false,
		ShowFullLevel:   false,
		HideKeys:        true,
	})
	logger.Driver = driver
	return logger
}

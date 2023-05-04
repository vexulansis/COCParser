package main

import (
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type MainLogger struct {
	Logger *logrus.Logger
}
type MainLoggerFields struct {
	Source      string
	Method      string
	Subject     string
	Destination string
}

// Example: [WORKER] [INSERT] [DATA] [CLANS] ..
func InitMainLogger() *MainLogger {
	mainLogger := &MainLogger{}
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&formatter.Formatter{
		FieldsOrder:     []string{"source", "method", "subject", "destination"},
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        false,
		ShowFullLevel:   false,
		HideKeys:        true,
	})
	mainLogger.Logger = logger
	return mainLogger
}
func (l *MainLogger) Print(f MainLoggerFields, data any) {
	fields := convertFields(f)
	switch extra := data.(type) {
	case int:
		l.Logger.WithFields(fields).Info()
	case string:
		l.Logger.WithFields(fields).Info(extra)
	case error:
		l.Logger.WithFields(fields).Error(extra)
	}
}
func (l *MainLogger) Fatal(f MainLoggerFields, data any) {
	fields := convertFields(f)
	l.Logger.WithFields(fields).Fatal(data)
}
func convertFields(f MainLoggerFields) logrus.Fields {
	fields := make(logrus.Fields)
	if len(f.Source) > 0 {
		fields["source"] = f.Source
	}
	if len(f.Method) > 0 {
		fields["method"] = f.Method
	}
	if len(f.Subject) > 0 {
		fields["subject"] = f.Subject
	}
	if len(f.Destination) > 0 {
		fields["destination"] = f.Destination
	}
	return fields

}

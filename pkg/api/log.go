package api

import (
	"fmt"
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type APILogger struct {
	Logger *logrus.Logger
}
type APILoggerFields struct {
	Source      string
	Method      string
	Subject     string
	Destination string
}

// Example: [WORKER] [GET] [KEY] [/api/apikey/list] ..
func initAPILogger() *APILogger {
	apiLogger := &APILogger{}
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
	apiLogger.Logger = logger
	return apiLogger
}
func (l *APILogger) Print(f APILoggerFields, data any) {
	fields := convertFields(f)
	switch extra := data.(type) {
	case *resty.Response:
		logstr := fmt.Sprintf("Status Code: %d", extra.StatusCode())
		switch {
		case extra.StatusCode() < 300:
			l.Logger.WithFields(fields).Info(logstr)
		case extra.StatusCode() >= 400:
			l.Logger.WithFields(fields).Error(logstr)
		default:
			l.Logger.WithFields(fields).Warn(logstr)
		}
	case int:
		l.Logger.WithFields(fields).Info()
	case string:
		l.Logger.WithFields(fields).Info(extra)
	case error:
		l.Logger.WithFields(fields).Error(extra)
	}

}
func convertFields(f APILoggerFields) logrus.Fields {
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

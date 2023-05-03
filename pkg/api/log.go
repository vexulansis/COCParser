package api

import (
	"fmt"
	"net/http"
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type APILogger struct {
	Logger *logrus.Logger
}
type APILoggerFields struct {
	Source   string
	Method   string
	Endpoint string
}

// Example: [WORKER] [GET] [/api/login] ..
func initAPILogger() *APILogger {
	apiLogger := &APILogger{}
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
	apiLogger.Logger = logger
	return apiLogger
}

func (l *APILogger) Print(f APILoggerFields, r *resty.Response) {
	logstr := fmt.Sprintf("Status: %d", r.StatusCode())
	fields := convertFields(f)
	switch r.StatusCode() {
	case http.StatusOK:
		l.Logger.WithFields(fields).Info(logstr)
	case http.StatusNotFound:
		l.Logger.WithFields(fields).Error(logstr)
	default:
		l.Logger.WithFields(fields).Warn(logstr)
	}

}
func convertFields(f APILoggerFields) logrus.Fields {
	fields := make(logrus.Fields)
	fields["source"] = f.Source
	fields["method"] = f.Method
	fields["endpoint"] = f.Endpoint
	return fields

}

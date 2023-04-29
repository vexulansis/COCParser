package api

import (
	"fmt"
	"net/http"
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type HTTPLogger struct {
	Logger *logrus.Logger
}
type HTTPFields struct {
	Source   string
	Method   string
	Endpoint string
}

func initHTTPLogger() *HTTPLogger {
	httplog := new(HTTPLogger)
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
	httplog.Logger = logger
	return httplog
}
func (l *HTTPLogger) Do(h HTTPFields, r *resty.Response) {
	logstr := fmt.Sprintf("Status: %d", r.StatusCode())
	switch r.StatusCode() {
	case http.StatusOK:
		l.Logger.WithFields(h.createHTTPFields()).Info(logstr)
	case http.StatusNotFound:
		l.Logger.WithFields(h.createHTTPFields()).Error(logstr)
	default:
		l.Logger.WithFields(h.createHTTPFields()).Warn(logstr)
	}
}
func (h *HTTPFields) createHTTPFields() logrus.Fields {
	f := make(logrus.Fields)
	f["source"] = h.Source
	f["method"] = h.Method
	f["endpoint"] = h.Endpoint
	return f
}

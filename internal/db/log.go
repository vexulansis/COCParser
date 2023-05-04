package db

import (
	"database/sql"
	"fmt"
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type DBLogger struct {
	Logger *logrus.Logger
}
type DBLoggerFields struct {
	Source      string
	Method      string
	Subject     string
	Destination string
}

// Example: [WORKER] [INSERT] [DATA] [CLANS] ..
func initDBLogger() *DBLogger {
	dbLogger := &DBLogger{}
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
	dbLogger.Logger = logger
	return dbLogger
}
func (l *DBLogger) Print(f DBLoggerFields, data any) {
	fields := convertFields(f)
	switch extra := data.(type) {
	case sql.Result:
		rows, err := extra.RowsAffected()
		if err != nil {
			l.Logger.WithFields(fields).Error(err)
		}
		logstr := fmt.Sprintf("Rows Affected: %d", rows)
		switch {
		case rows == 0:
			l.Logger.WithFields(fields).Warn(logstr)
		default:
			l.Logger.WithFields(fields).Info(logstr)
		}
	case int:
		l.Logger.WithFields(fields).Info()
	case string:
		l.Logger.WithFields(fields).Info(extra)
	case error:
		l.Logger.WithFields(fields).Error(extra)
	}

}
func convertFields(f DBLoggerFields) logrus.Fields {
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

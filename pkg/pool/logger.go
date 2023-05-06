package pool

import (
	"os"
	"strconv"
	"strings"

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

func (l *Logger) Info(log string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Info()
}
func (l *Logger) InfoF(log string, extra string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Info(extra)
}
func (l *Logger) Warn(log string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Warn()
}
func (l *Logger) WarnF(log string, extra string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Warn(extra)
}
func (l *Logger) Error(log string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Error()
}
func (l *Logger) ErrorF(log string, extra string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Error(extra)
}
func (l *Logger) Fatal(log string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Fatal()
}
func (l *Logger) FatalF(log string, extra string) {
	fields := makeFields(log)
	l.Driver.WithFields(fields).Fatal(extra)
}
func makeFields(log string) logrus.Fields {
	fields := make(logrus.Fields)
	logArr := strings.Fields(log)
	for i, f := range logArr {
		fields[strconv.Itoa(i)] = f
	}
	return fields
}

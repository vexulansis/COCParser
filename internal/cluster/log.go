package cluster

import (
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type ClusterLogger struct {
	Logger *logrus.Logger
}

// Example: [WORKER] [PRODUCE] [TAG] [TOPIC] ..
func initClusterLogger() *ClusterLogger {
	clusterlogger := &ClusterLogger{}
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&formatter.Formatter{
		FieldsOrder:     []string{"source", "method", "subject", "destination"},
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        false,
		ShowFullLevel:   true,
		HideKeys:        true,
	})
	clusterlogger.Logger = logger
	return clusterlogger
}

package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	fmt.Println("runned")
	logrus.SetOutput(os.Stdout)
	logrus.Info("from logrus init")
}*/

// LOGGER global instances
var LOGGER = &logrus.Logger{
	Out: os.Stdout,
	Formatter: &logrus.TextFormatter{
		FullTimestamp: true,
	},
}

// Get get Logger instance from logrus
func Get() *logrus.Logger {
	return &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
		},
	}
}

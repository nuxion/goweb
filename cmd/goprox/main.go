package main

import (
	"os"

	"github.com/nuxion/goweb/pkg/proxy"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Info("MAIN")
	proxy.Proxy()
}

package main

import (
	"flag"
	"os"

	"github.com/nuxion/goweb/pkg/config"
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
	pathConf := flag.String("cfg", "config.example.toml", "fullpath to config.toml file")
	listenPort := flag.String("port", "9090", "port where the service will be listening")

	flag.Parse()
	//var path string = *pathConf
	conf, _ := config.LoadTom(*pathConf)
	conf.Port = *listenPort
	logrus.Info("Ready&Go Executing...")
	logrus.Debug("Using ", *pathConf)
	proxy.Run(conf)
}

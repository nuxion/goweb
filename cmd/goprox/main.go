package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nuxion/goweb/pkg/config"
	"github.com/nuxion/goweb/pkg/proxy"
	"github.com/sirupsen/logrus"
)

// Version tag commit label
var Version = "dev"

// Build  hash of the commit built
var Build = "init"

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
	version := flag.Bool("version", false, "print version")

	flag.Parse()
	if *version == true {
		fmt.Println("Version: ", Version)
		fmt.Println("Build: ", Build)
		os.Exit(0)

	}
	//var path string = *pathConf
	conf, _ := config.LoadTom(*pathConf)
	conf.Port = *listenPort
	logrus.Info("Ready&Go Executing...")
	logrus.Debug("Using ", *pathConf)
	proxy.Run(conf)
}

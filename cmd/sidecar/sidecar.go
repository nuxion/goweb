package main

import (
	"flag"
	"github.com/nuxion/goweb/pkg/httpserver"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "Listen port")
	flag.Parse()
	httpserver.Run(port)
}

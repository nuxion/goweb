package main

import (
	"fmt"
	"net/http"
	"flag"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// Index is a simple middleware
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Infof("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.SetOutput(os.Stdout)
	var port string
	flag.StringVar(&port, "port", "8080", "Listen port")
	flag.Parse()
	log.Info("Starting at... :", port)
	finalPort := fmt.Sprintf(":%s", port)
		log.Fatal(http.ListenAndServe(finalPort, logHandler(router)))
}

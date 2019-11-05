package main

import (
	"fmt"
	"net/http"
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
		log.Infof("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.SetOutput(os.Stdout)
	log.Info("Starting at... :8080")
	log.Fatal(http.ListenAndServe(":8080", logHandler(router)))
}

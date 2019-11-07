package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httputil"
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
		dump, _ := httputil.DumpRequest(r, true)
		log.Info(string(dump))
		handler.ServeHTTP(w, r)
	})
}

// Run is the exported instance
func Run(port string) {
	router := httprouter.New()
	router.GET("/", Index)
	log.SetOutput(os.Stdout)
	log.Info("Starting at... :", port)
	finalPort := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(finalPort, logHandler(router)))
}

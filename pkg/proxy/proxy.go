package proxy

import (
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nuxion/goweb/pkg/config"
	limiter "github.com/nuxion/goweb/pkg/ratelimiter"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

// ClientID struct
type ClientID struct {
	id    []byte
	ip    string
	agent string
	hash  string
}

// NewClient creates a new ClientID object, based on the RemoteIP and
// the UserAgent
func NewClient(i, a string) *ClientID {
	id := []byte(i)
	id = append(id, a...)
	c := &ClientID{id: id, ip: i, agent: a}
	c.Hash()

	return c
}

// Hash hashing clientid
func (c *ClientID) Hash() {
	h := sha1.New()
	h.Write(c.id)
	c.hash = base64.URLEncoding.EncodeToString(h.Sum(nil))
	log.Debug("hash from Hash() ", c.hash)
	//return &c.hash
}

// NewMultipleHostReverseProxy Multihost proxy
func NewMultipleHostReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(r *http.Request) {

		target := targets[rand.Int()%len(targets)]
		guid := xid.New()
		r.Header.Add("X-Request-ID", guid.String())
		r.URL.Scheme = target.Scheme
		r.URL.Host = target.Host
		log.Info("RemoteAddr: ", r.RemoteAddr)
		dump, _ := httputil.DumpRequest(r, true)
		log.Info(string(dump))
		cIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Debug("RemoteIP: ", cIP)
		client := NewClient(cIP, r.UserAgent())
		log.Debug("ClientID hash: ", client.hash)

	}
	return &httputil.ReverseProxy{Director: director}

}

// prepareUrls by service
func prepareUrls(s *config.Service) []*url.URL {
	// var urls [len(s.Hosts)]*url.URL
	lenght := len(s.Hosts)
	urls := make([]*url.URL, lenght)
	for i, e := range s.Hosts {
		urls[i] = &url.URL{Scheme: s.Proto, Host: e}
	}

	return urls

}

// Proxy main function to run proxy
func Proxy() {
	proxy := NewMultipleHostReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:8082",
		},
		{
			Scheme: "http",
			Host:   "localhost:8081",
		},
	})
	log.Fatal(http.ListenAndServe(":9090", proxy))
}

func logH(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		h.ServeHTTP(w, r) // call original
		log.Println("After")
	})
}

// Run new main function to run proxy
func Run(c *config.Config) {
	service := c.Services["httpserver"]
	u := prepareUrls(&service)
	proxy := NewMultipleHostReverseProxy(u)
	http.Handle("/", limiter.SimpleLimiter(proxy))
	//handler := ratelimiter.SimpleLimiter(proxy)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

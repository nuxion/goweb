package proxy

import (
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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
		/*r.Header.Add("X-Forwarded-Host", r.Host)
		r.Header.Add("X-Origin-Host", target.Host)*/
		guid := xid.New()
		r.Header.Add("X-Request-ID", guid.String())
		r.URL.Scheme = target.Scheme
		r.URL.Host = target.Host
		r.URL.Path = target.Path
		log.Info("RemoteAddr: ", r.RemoteAddr)
		dump, _ := httputil.DumpRequest(r, true)
		log.Info(string(dump))
		// FIXME only works for IPV4
		cIP := strings.Split(r.RemoteAddr, ":")[0]
		log.Debug("RemoteIP: ", cIP)
		client := NewClient(cIP, r.UserAgent())
		log.Debug("ClientID hash: ", client.hash)

	}
	return &httputil.ReverseProxy{Director: director}

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

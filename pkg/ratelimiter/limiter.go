package ratelimiter

import (
	"net"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

type LimitContext struct {
	rLimit rate.Limit
	burst  int
}

func (lc *LimitContext) addVisitor(ip string) *rate.Limiter {
	limiter := rate.NewLimiter(lc.rLimit, lc.burst)
	//mu.Lock()
	// Include the current time when creating a new visitor.
	visitors[ip] = &visitor{limiter, time.Now()}
	//mu.Unlock()
	return limiter
}

func NewLimitContext(r rate.Limit, b int) *LimitContext {
	return &LimitContext{r, b}
}

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Change the the map to hold values of the type visitor.
var visitors = make(map[string]*visitor)
var mu sync.Mutex

// Run a background goroutine to remove old entries from the visitors map.
/*func init() {
	go CleanupVisitors()
}*/

func (lc *LimitContext) getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		return lc.addVisitor(ip)
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	return v.limiter
}

// CleanupVisitors Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func CleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		defer mu.Unlock()
		for ip, v := range visitors {
			if time.Now().Sub(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
	}
}

// SimpleLimiter a lot of improve needed
func (lc *LimitContext) SimpleLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		log.Info("Inside SimpleLimiter ", r.RemoteAddr)
		if err != nil {
			log.Info(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := lc.getVisitor(ip)
		log.Info("after visitor")
		log.Info(limiter.Limit())
		if limiter.Allow() == false {
			log.Error("Not Allowed")
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

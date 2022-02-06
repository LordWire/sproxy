package main

// this will handle the actual request.
// All processing (e.g. header rewrites) logic should be implemented here.

// If we want correct metrics, this function should be broken to whatever
// processing happens inside the proxy, and then anotherone that just execs a ServeHTTP.
// This will allow for correct instrumentation and separation between internal processing
// and actual backend delays/status

// in an extended version of this program, this structure would allow wrapper functions
// that choose either the instrumented or the non-instrumented versions of the actual
// handlers, in order to be able to turn monitoring on or off.

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// wrapper function to be used instead of serveReverseProxy
func instrumentedReverseProxy(pattern string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next(w, r)
		handlerDuration.WithLabelValues(pattern).Observe(time.Since(now).Seconds())
	})
}

// the actual function that handles reverse proxy logic.
func serveReverseProxy(res http.ResponseWriter, req *http.Request) {
	requestsInProcess.Inc()
	backendUrl, err := url.Parse(backendList.getRandomBackend())
	log.Println("serving request from proxy address:" + backendUrl.String())

	if err != nil {
		panic("error parsing url")
	}

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(backendUrl)
	proxy.ServeHTTP(res, req)
	requestsInProcess.Dec()
}

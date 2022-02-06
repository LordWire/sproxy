package main

// Implementation of metrics.
// Currently we measure the total connection requests and
// the total request serving time.

import (
	"github.com/prometheus/client_golang/prometheus"
)

var requestsInProcess = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "requests_in_process",
		Help: "Current number of requests in-process",
	},
)

var (
	handlerDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "sproxy_total_request_duration_seconds",
		Help: "Total turnaround time for a request.",
	}, []string{"path"})
)

//register our prometheus metrics
func registerMetrics() {
	prometheus.MustRegister(requestsInProcess)
	prometheus.MustRegister(handlerDuration)
}

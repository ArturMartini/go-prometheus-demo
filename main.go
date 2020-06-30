package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "myapp_processed_ops_total",
		Help:        "The total number of processed events",
		ConstLabels: nil,
	})
)

func main() {
	http.HandleFunc("/demo", countDemo)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

func countDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("count")
	opsProcessed.Inc()
	w.WriteHeader(http.StatusOK)
}

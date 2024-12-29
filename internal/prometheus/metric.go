package metrics

import (
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var GoroutinesMetric = prometheus.NewGaugeFunc(
	prometheus.GaugeOpts{
		Name: "num_goroutines",
		Help: "Current number of goroutines",
	},
	func() float64 {
		return float64(runtime.NumGoroutine())
	},
)

func InitMetrics(port string) {
	prometheus.MustRegister(GoroutinesMetric)
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		logrus.Printf("Starting metrics server on port %s\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			logrus.Fatalf("Failed to start metrics server: %v", err)
		}
	}()
}

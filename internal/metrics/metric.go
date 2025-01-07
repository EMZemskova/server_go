package metrics

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
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

var (
	HttpStatusMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_responses_total",
			Help: "Count of HTTP responses, labeled by status code and method",
		},
		[]string{"status", "method"},
	)
)

func InitMetrics(port string) {
	prometheus.MustRegister(GoroutinesMetric)
	prometheus.MustRegister(HttpStatusMetric)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		logrus.Printf("Starting metrics server on port %s\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			logrus.Fatalf("Failed to start metrics server: %v", err)
		}
	}()
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		statusCode := c.Writer.Status()
		method := c.Request.Method
		HttpStatusMetric.WithLabelValues(http.StatusText(statusCode), method).Inc()
	}
}

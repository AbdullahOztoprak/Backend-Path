package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// Define metrics
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "route"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)
)

// InitMetrics initializes the metrics and registers them with Prometheus
func InitMetrics() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

// Middleware for tracking metrics
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := r.URL.Path
		method := r.Method

		// Start timer
		timer := prometheus.NewTimer(httpRequestDuration.WithLabelValues(method, route))
		defer timer.ObserveDuration()

		// Increment the counter
		httpRequestsTotal.WithLabelValues(method, route).Inc()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Handler for exposing metrics
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
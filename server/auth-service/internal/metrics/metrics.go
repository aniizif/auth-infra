package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		}, []string{"method", "endpoint", "status"},
	)

	HTTPRequestsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_requests_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
		}, []string{"method", "endpoint"})

	AUTHSuccessTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_success_total",
			Help: "Total number of Success authorization",
		}, []string{"endpoint"},
	)

	AUTHFailureTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_failure_total",
			Help: "Total number of Failure authorization",
		}, []string{"endpoint"},
	)

	DBQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "db_query_duration_seconds",
			Help: "Duration of DB Query in seconds",
		}, []string{"operation", "table"},
	)
)

func init() {
	prometheus.MustRegister(HTTPRequestsTotal, HTTPRequestsDuration, AUTHSuccessTotal, AUTHFailureTotal, DBQueryDuration)
}

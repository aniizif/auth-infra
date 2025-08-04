package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"

	"github.com/aniizif/stack-mate/auth-service/internal/metrics"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		endpoint := c.FullPath()

		if endpoint == "" {
			endpoint = "unknown"
		}

		status := strconv.Itoa(c.Writer.Status())

		metrics.HTTPRequestsTotal.WithLabelValues(c.Request.Method, endpoint, status).Inc()
		metrics.HTTPRequestsDuration.WithLabelValues(c.Request.Method, endpoint).Observe(duration)
	}
}

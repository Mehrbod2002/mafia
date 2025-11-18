package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger attaches simple request logging to gin.Engine.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"latency": latency.String(),
		}).Info("request completed")
	}
}

// CorrelationID ensures each request has a correlation id header.
func CorrelationID(header string) gin.HandlerFunc {
	if header == "" {
		header = "X-Correlation-ID"
	}
	return func(c *gin.Context) {
		if c.GetHeader(header) == "" {
			c.Request.Header.Set(header, time.Now().Format("20060102150405.000"))
		}
		c.Writer.Header().Set(header, c.GetHeader(header))
		c.Next()
	}
}

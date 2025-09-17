package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func SlowRequestLogger(threshold time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		if latency > threshold {
			log.Printf("SLOW REQUEST: [%s] %s - %d (%v)",
				c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency)
		}

	}

}

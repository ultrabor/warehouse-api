package v1

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		elapsed := time.Since(start)

		logger.Info("http request",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.RawPath),
			slog.Int("status", c.Writer.Status()),
			slog.Duration("duration", elapsed))
	}
}

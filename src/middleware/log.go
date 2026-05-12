package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(log *zap.Logger, skipIPs ...string) gin.HandlerFunc {
	skipSet := make(map[string]struct{}, len(skipIPs))
	for _, ip := range skipIPs {
		skipSet[ip] = struct{}{}
	}

	return func(c *gin.Context) {

		start := time.Now()

		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		if _, blocked := skipSet[clientIP]; blocked {
			c.Next()
			return
		}

		traceID, _ := c.Get(TraceIDKey)
		traceIDStr, _ := traceID.(string)

		c.Next()
		latency := time.Since(start)

		statusCode := c.Writer.Status()

		fields := []zap.Field{
			zap.String("trace_id", traceIDStr),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", statusCode),
			zap.String("client_ip", clientIP),
			zap.String("user_agent", userAgent),
			zap.Int64("latency_ms", latency.Milliseconds()),
		}

		if len(c.Errors) > 0 {
			fields = append(fields,
				zap.String("error", c.Errors.String()),
			)

			log.Error("request failed", fields...)
			return
		}

		if statusCode >= 500 {
			log.Error("server error", fields...)
			return
		}

		log.Info("request completed", fields...)
	}
}

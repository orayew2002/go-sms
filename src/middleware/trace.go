package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const TraceIDKey = "trace_id"

func TraceID() gin.HandlerFunc {
	return func(c *gin.Context) {

		traceID := c.GetHeader("X-Trace-ID")

		if traceID == "" {
			traceID = uuid.NewString()
		}

		c.Set(TraceIDKey, traceID)
		c.Writer.Header().Set("X-Trace-ID", traceID)

		c.Next()
	}
}

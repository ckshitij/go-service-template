package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	REQ_ID = "request_id"
)

func AttachRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set(REQ_ID, requestID)
		c.Next()
	}
}

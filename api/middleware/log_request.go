package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Headers: %v", c.Request.Header)
		log.Printf("Body: %v", c.Request.Body)
		requestId := uuid.New().String()
		c.Set("request_id", requestId)
		c.Next()
	}
}

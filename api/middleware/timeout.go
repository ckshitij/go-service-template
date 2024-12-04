package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeoutMiddleware adds a context with timeout to each request
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel() // Ensure the context is canceled when the request is finished

		// Replace the request's context with the new context
		c.Request = c.Request.WithContext(ctx)

		// Proceed with the request
		c.Next()

	}
}

package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.io/ckshitij/go-service-template/config"
)

// SetupServer initializes and configures the Gin server
func SetupServer(conf *config.Config) *gin.Engine {
	r := gin.Default()

	// Example route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// Add more routes here

	return r
}

// StartServer starts the Gin server
func StartServer(conf *config.Config) error {
	r := SetupServer(conf)
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	return r.Run(address)
}

package server

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.io/ckshitij/go-service-template/api/middleware"
	"github.io/ckshitij/go-service-template/api/pkg/users"
	"github.io/ckshitij/go-service-template/api/wrapper/rest"
	"github.io/ckshitij/go-service-template/config"
	"github.io/ckshitij/go-service-template/db"
)

// SetupServer initializes and configures the Gin server
func SetupServer(conf *config.Config) *gin.Engine {
	router := gin.Default()

	conn, err := db.NewPostgresDB(conf)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router.Use(middleware.TimeoutMiddleware(3 * time.Second))
	usersHandlers := users.InitUsers(conn)
	endpoints := []rest.IEndpointProvider{
		usersHandlers,
	}
	rest.RegisterHandlers(router, endpoints)

	return router
}

// StartServer starts the Gin server
func StartServer(conf *config.Config) error {
	r := SetupServer(conf)
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	return r.Run(address)
}

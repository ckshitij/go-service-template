package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.io/ckshitij/go-service-template/api/pkg/users"
	"github.io/ckshitij/go-service-template/api/wrapper/rest"
	"github.io/ckshitij/go-service-template/config"
	"github.io/ckshitij/go-service-template/db"
)

// SetupServer initializes and configures the Gin server
func SetupServer(conf *config.Config) *gin.Engine {
	r := gin.Default()

	conn, err := db.NewPostgresDB(conf)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	usersHandlers := users.InitUsers(conn).Handlers()
	rest.ProcessAndRegisterHandlers(r, usersHandlers)

	return r
}

// StartServer starts the Gin server
func StartServer(conf *config.Config) error {
	r := SetupServer(conf)
	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	return r.Run(address)
}

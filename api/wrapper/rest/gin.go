package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IEndpointProvider interface {
	Handlers() []HTTPHandler
}

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

type HTTPHandler struct {
	Method     HTTPMethod
	Middleware []gin.HandlerFunc
	Handler    gin.HandlerFunc
	Path       string
}

func RegisterHandlers(router *gin.Engine, providers []IEndpointProvider) {

	grp := router.Group("/template-service/api/v1/")

	for _, provider := range providers {
		handlers := provider.Handlers()
		for _, handler := range handlers {
			grp.Handle(string(handler.Method), handler.Path, append(handler.Middleware, handler.Handler)...)
		}
	}

	docPath := grp.BasePath() + "doc"
	// Serve the Swagger YAML file
	router.GET(docPath, func(c *gin.Context) {
		c.File("swagger/docs/swagger.yaml")
	})

	// Serve the Swagger UI
	// Serve the Swagger UI and static files
	grp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(docPath)))
}

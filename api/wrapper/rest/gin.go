package rest

import "github.com/gin-gonic/gin"

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

func ProcessAndRegisterHandlers(router *gin.Engine, handlers []HTTPHandler) {
	for _, handler := range handlers {
		router.Handle(string(handler.Method), handler.Path, append(handler.Middleware, handler.Handler)...)
	}
}

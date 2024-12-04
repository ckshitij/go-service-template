package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.io/ckshitij/go-service-template/api/middleware"
	"github.io/ckshitij/go-service-template/api/wrapper/rest"
)

type usersHandler struct {
	service UsersService
}

func NewUsersHandler(service UsersService) *usersHandler {
	return &usersHandler{service: service}
}

func (h *usersHandler) Handlers() []rest.HTTPHandler {
	return []rest.HTTPHandler{
		{
			Method:     rest.POST,
			Path:       "/users",
			Middleware: []gin.HandlerFunc{middleware.AttachRequestID()},
			Handler:    h.CreateUser,
		},
		{
			Method:  rest.GET,
			Path:    "/users",
			Handler: h.GetUserByEmail,
		},
	}
}

func (h *usersHandler) CreateUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(c.Request.Context(), ServiceRequest{
		UserRequest: req,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *usersHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	user, err := h.service.GetUser(c.Request.Context(), ServiceRequest{
		UserRequest: UserRequest{
			Email:    email,
			Password: password,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

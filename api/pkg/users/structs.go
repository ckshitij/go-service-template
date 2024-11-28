package users

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// User represents the user entity
type User struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"-"` // "-" means this field won't be included in JSON
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *time.Time      `json:"deleted_at,omitempty"`
}

// CreateUserRequest represents the request body for creating a user
type UserRequest struct {
	Name     string          `json:"name" binding:"required"`
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required,min=6"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}

type ServiceRequest struct {
	UserRequest
}

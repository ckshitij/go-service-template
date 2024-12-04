package users

import (
	"database/sql"

	"github.io/ckshitij/go-service-template/api/wrapper/rest"
)

func InitUsers(db *sql.DB) rest.IEndpointProvider {
	repo := NewUsersRepository(db)
	service := NewUsersService(repo)
	return NewUsersHandler(service)
}

package users

import "database/sql"

func InitUsers(db *sql.DB) *usersHandler {
	repo := NewUsersRepository(db)
	service := NewUsersService(repo)
	return NewUsersHandler(service)
}

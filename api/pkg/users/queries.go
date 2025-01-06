package users

const (
	GetUsers = `
		SELECT 
			id, name, email, password, metadata, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1
	`

	CreateUser = `
		INSERT INTO users (name, email, password, metadata)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
)

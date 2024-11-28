package users

import (
	"context"
	"database/sql"
)

type usersRepoImpl struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepoImpl{db: db}
}

func (r *usersRepoImpl) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (name, email, password, metadata)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Metadata,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *usersRepoImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, name, email, password, metadata, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Metadata,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	return user, err
}

// Implement other repository methods...

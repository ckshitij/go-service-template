package users

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) error {
	return r.db.QueryRowContext(
		ctx,
		CreateUser,
		user.Name,
		user.Email,
		user.Password,
		user.Metadata,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	err := r.db.QueryRowContext(ctx, GetUsers, email).Scan(
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

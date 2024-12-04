package users

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type usersServiceImpl struct {
	repo UsersRepository
}

func NewUsersService(repo UsersRepository) UsersService {
	return &usersServiceImpl{repo: repo}
}

func (s *usersServiceImpl) CreateUser(ctx context.Context, req ServiceRequest) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Metadata: req.Metadata,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersServiceImpl) GetUser(ctx context.Context, req ServiceRequest) (*User, error) {

	userInfo, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password)); err != nil {
		return nil, err
	}
	return userInfo, nil
}

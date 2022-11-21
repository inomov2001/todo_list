package service

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
func (s Service) RegisterUser(ctx context.Context, name string) (int, error) {
	return s.repo.CreateUser(ctx, name)
}

func (s Service) UpdateUser(ctx context.Context, user User) error {
	return s.repo.UpdateUser(ctx, user)
}

func (s Service) DeleteUser(ctx context.Context, user User) error {
	return s.repo.DeleteUser(ctx, user)
}

func (s Service) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.ListUsers(ctx)
}

func (s Service) GetUserBYID(ctx context.Context, id int) (User, error) {
	return s.repo.GetUserBYID(ctx, id)
}

type Repository interface {
	CreateUser(ctx context.Context, name string) (int, error)
	ListUsers(ctx context.Context) ([]User, error)
	GetUserBYID(ctx context.Context, id int) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, user User) error
}

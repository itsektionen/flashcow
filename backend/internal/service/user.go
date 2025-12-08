package service

import (
	"context"
	"log/slog"

	"github.com/itsektionen/flashcow/backend/internal/model"
	"github.com/itsektionen/flashcow/backend/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	ListUsers(ctx context.Context) ([]model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	// TODO: Make sure it's valid user details

	slog.Info("Creating user", user.FullName, user.ChapterEmailAddress)

	if err := s.repository.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	// TODO: Authentication?

	return s.repository.GetUser(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context) ([]model.User, error) {
	return s.repository.ListUsers(ctx)
}

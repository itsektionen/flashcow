package service

import (
	"context"
	"log/slog"

	"github.com/itsektionen/flashcow/backend/internal/model"
	"github.com/itsektionen/flashcow/backend/internal/repository"
)

type UserService struct {
	Repository *repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	// TODO: Make sure it's valid user details

	slog.Info("Creating user", user.FullName, user.ChapterEmailAddress)

	if err := s.Repository.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	// TODO: Authentication?

	return s.Repository.GetUser(ctx, id)
}

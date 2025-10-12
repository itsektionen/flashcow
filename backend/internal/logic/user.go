package logic

import (
	"log/slog"

	"github.com/itsektionen/flashcow/backend/internal"
)

func (s *Service) CreateUser(user internal.User) (*internal.User, error) {
	// TODO: Make sure it's valid user details

	slog.Info("Creating user", user.FullName, user.ChapterEmailAddress)

	if err := s.Repo.CreateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

package service

import (
	"context"

	"github.com/itsektionen/flashcow/backend/internal/model"
	"github.com/itsektionen/flashcow/backend/internal/repository"
)

type CommitteeService struct {
	Repository *repository.CommitteeRepository
}

func (s *CommitteeService) ListCommittees(ctx context.Context) ([]model.Committee, error) {
	committees, err := s.Repository.ListCommittees(ctx)

	if err != nil {
		return nil, err
	}

	return committees, nil
}

func (s *CommitteeService) GetCommittee(ctx context.Context, id int64) (*model.Committee, error) {
	return s.Repository.GetCommittee(ctx, id)
}

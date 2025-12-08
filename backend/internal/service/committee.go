package service

import (
	"context"

	"github.com/itsektionen/flashcow/backend/internal/model"
	"github.com/itsektionen/flashcow/backend/internal/repository"
)

type CommitteeService interface {
	ListCommittees(ctx context.Context) ([]model.Committee, error)
	GetCommittee(ctx context.Context, id int64) (*model.Committee, error)
}

type committeeService struct {
	repository repository.CommitteeRepository
}

func NewCommitteeService(repository repository.CommitteeRepository) CommitteeService {
	return &committeeService{
		repository: repository,
	}
}

func (s *committeeService) ListCommittees(ctx context.Context) ([]model.Committee, error) {
	committees, err := s.repository.ListCommittees(ctx)

	if err != nil {
		return nil, err
	}

	return committees, nil
}

func (s *committeeService) GetCommittee(ctx context.Context, id int64) (*model.Committee, error) {
	return s.repository.GetCommittee(ctx, id)
}

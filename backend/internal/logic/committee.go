package logic

import (
	"context"

	"github.com/itsektionen/flashcow/backend/internal"
	"github.com/itsektionen/flashcow/backend/internal/repository"
)

type CommitteeService struct {
	Database *repository.Database
}

func (s *CommitteeService) ListCommittees(ctx context.Context) ([]internal.Committee, error) {
	committees, err := s.Database.ListCommittees(ctx)

	if err != nil {
		return nil, err
	}

	return committees, nil
}

func (s *CommitteeService) GetCommittee(ctx context.Context, id int64) (*internal.Committee, error) {
	return s.Database.GetCommittee(ctx, id)
}

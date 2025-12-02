package logic

import (
	"context"

	"github.com/itsektionen/flashcow/backend/internal/storage"
)

type CommitteeService struct {
	Database *storage.Database
}

func (s *CommitteeService) GetCommittees(ctx context.Context) ([][]string, error) {
	committees, err := s.Database.GetCommittees(ctx)

	if err != nil {
		return nil, err
	}

	return committees, nil
}

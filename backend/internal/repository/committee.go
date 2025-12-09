package repository

import (
	"context"
	"slices"

	"github.com/itsektionen/flashcow/backend/internal/model"
)

type CommitteeRepository interface {
	ListCommittees(ctx context.Context) ([]model.Committee, error)
	GetCommittee(ctx context.Context, id int64) (*model.Committee, error)
}

type committeeRepository struct{}

func NewCommitteeRepository() CommitteeRepository {
	return &committeeRepository{}
}

// TODO: Add all committees or call our global datastore (if we have one)
var committees = []model.Committee{
	{ID: 1, ShortName: "ITK", Name: "ITerativa Klubben"},
	{ID: 2, ShortName: "QMISK", Name: "QlubbMästeriet IT-Sektionen Kista"},
	{ID: 3, ShortName: "TMEIT", Name: "TraditionsMEsterIT"},
	{ID: 4, ShortName: "SMN", Name: "Studiemiljönämnden"},
}

func (r *committeeRepository) ListCommittees(ctx context.Context) ([]model.Committee, error) {
	return committees, nil
}

func (r *committeeRepository) GetCommittee(ctx context.Context, id int64) (*model.Committee, error) {
	idx := slices.IndexFunc(committees, func(c model.Committee) bool {
		return c.ID == id
	})

	committee := &committees[idx]

	return committee, nil
}

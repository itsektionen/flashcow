package repository

import (
	"context"
	"slices"

	"github.com/itsektionen/flashcow/backend/internal"
)

// TODO: Add all committees or call our global datastore (if we have one)
var committees = []internal.Committee{
	{ID: 1, ShortName: "ITK", Name: "ITerativa Klubben"},
	{ID: 2, ShortName: "QMISK", Name: "QlubbMästeriet IT-Sektionen Kista"},
	{ID: 3, ShortName: "TMEIT", Name: "TraditionsMEsterIT"},
	{ID: 4, ShortName: "SMN", Name: "Studiemiljönämnden"},
}

func (db *Database) ListCommittees(ctx context.Context) ([]internal.Committee, error) {
	return committees, nil
}

func (db *Database) GetCommittee(ctx context.Context, id int64) (*internal.Committee, error) {
	idx := slices.IndexFunc(committees, func(c internal.Committee) bool {
		return c.ID == id
	})

	committee := &committees[idx]

	if committee == nil {
		return nil, ErrNotFound
	}

	return committee, nil
}

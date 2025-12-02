package storage

import "context"

// TODO: Add all committees or call our global datastore (if we have one)
var committees = [][]string{{"ITK", "ITerativa Klubben"}, {"QMISK", "Qlubbmästeriet IT-Sektionen Kista"}, {"TMEIT", "TraditionsMEsterIT"}, {"SMN", "Studiemiljönämnden"}}

func (db *Database) GetCommittees(ctx context.Context) ([][]string, error) {
	return committees, nil
}

package model

import "time"

type ReceiptReport struct {
	id             int64
	purchaseDate   time.Time
	submissionTime time.Time
	submitterId    int64
	committeeId    int64
	contents       string
	foodTotal      int64
	beerTotal      int64
	sodaTotal      int64
	ciderTotal     int64
	wineTotal      int64
	spiritsTotal   int64
	materialTotal  int64
	internRepTotal int64
	comments       string
	image          []byte
}

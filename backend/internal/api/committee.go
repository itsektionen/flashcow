package api

import "net/http"

type CommitteeHandler struct {
}

// TODO: Add all committees or call our global datastore (if we have one)
var committees = [][]string{{"ITK", "ITerativa Klubben"}, {"QMISK", "Qlubbmästeriet IT-Sektionen Kista"}, {"TMEIT", "TraditionsMEsterIT"}, {"SMN", "Studiemiljönämnden"}}

func (CommitteeHandler) getCommittees(w http.ResponseWriter, r *http.Request) {
	response(w, committees, http.StatusOK)
}

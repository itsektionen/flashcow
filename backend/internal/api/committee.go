package api

import "net/http"

type CommitteeHandler struct {
}

func (CommitteeHandler) getCommittees(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

package api

import (
	"net/http"

	"github.com/itsektionen/flashcow/backend/internal/logic"
)

type CommitteeHandler struct {
	CommitteeService *logic.CommitteeService
}

func (h *CommitteeHandler) getCommittees(w http.ResponseWriter, r *http.Request) {
	committees, err := h.CommitteeService.GetCommittees(r.Context())

	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusInternalServerError)
	}

	response(w, committees, http.StatusOK)
}

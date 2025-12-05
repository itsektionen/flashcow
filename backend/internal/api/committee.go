package api

import (
	"net/http"
	"strconv"

	"github.com/itsektionen/flashcow/backend/internal/service"
)

type CommitteeHandler struct {
	CommitteeService *service.CommitteeService
}

func (h *CommitteeHandler) listCommittees(w http.ResponseWriter, r *http.Request) {
	committees, err := h.CommitteeService.ListCommittees(r.Context())

	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusInternalServerError)
	}

	response(w, committees, http.StatusOK)
}

func (h *CommitteeHandler) getCommittee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}

	committee, err := h.CommitteeService.GetCommittee(r.Context(), id)

	if err != nil {
		errorResponse(w, err)
		return
	}

	response(w, committee, http.StatusOK)
}

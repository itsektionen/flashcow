package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/itsektionen/flashcow/backend/internal"
	"github.com/itsektionen/flashcow/backend/internal/logic"
)

type UserHandler struct {
	Service *logic.Service
}

func (UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	//id := r.PathValue("id")

	w.WriteHeader(http.StatusOK)
}

func (s *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	u := &internal.User{}

	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}

	user, err := s.Service.CreateUser(*u)

	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}

	response(w, user, http.StatusOK)
}

func (UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}
	u := &internal.User{}

	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}

	u.ID = id

	response(w, u, http.StatusOK)
}

func (UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	//id := r.PathValue("id")

	w.WriteHeader(http.StatusOK)
}

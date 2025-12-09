package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/itsektionen/flashcow/backend/internal/model"
	"github.com/itsektionen/flashcow/backend/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func (h *UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.ListUsers(r.Context())

	if err != nil {
		errorResponse(w, err)
		return
	}

	response(w, users, http.StatusOK)
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(r.Context(), id)

	if err != nil {
		errorResponse(w, err)
		return
	}

	response(w, user, http.StatusOK)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	u := &model.User{}

	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), *u)
	if err != nil {
		errorResponse(w, err)

		return
	}

	response(w, user, http.StatusOK)
}

func (*UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}
	u := &model.User{}

	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}

	u.ID = id

	response(w, u, http.StatusOK)
}

func (*UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	//id := r.PathValue("id")

	w.WriteHeader(http.StatusOK)
}

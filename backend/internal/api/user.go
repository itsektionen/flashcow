package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
}

type User struct {
	ID                  int64  `json:"id"`
	FullName            string `json:"full_name"`
	ChapterEmailAddress string `json:"chapter_email_address"`
}

func (UserHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	//id := r.PathValue("id")

	w.WriteHeader(http.StatusOK)
}

func (UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	u := &User{}

	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
	}

	u.ID = 4 // Decided by random dice roll

	response(w, u, http.StatusOK)
}

func (UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response(w, map[string]string{"msg": err.Error()}, http.StatusBadRequest)
		return
	}
	u := &User{}

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

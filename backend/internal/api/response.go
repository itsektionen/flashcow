package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/itsektionen/flashcow/backend/internal/repository"
)

func response(w http.ResponseWriter, data any, status int) {
	if data == nil {
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Could not write response")
		return
	}
}

type httpError struct {
	Message string `json:"msg"`
}

func errorResponse(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	switch {
	case errors.Is(err, repository.ErrNotFound):
		statusCode = http.StatusNotFound
	}

	slog.Error("error serving request", "err", err, "status_code", statusCode)

	e := &httpError{Message: err.Error()}

	response(w, e, statusCode)
}

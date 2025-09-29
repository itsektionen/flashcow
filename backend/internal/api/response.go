package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func response(w http.ResponseWriter, data any, status int) {
	if response == nil {
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

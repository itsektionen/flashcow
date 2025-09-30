package main

import (
	"log/slog"
	"os"

	"github.com/itsektionen/flashcow/backend/internal/api"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	s := &api.Server{
		Addr: ":8080",
	}

	if err := s.Serve(); err != nil {
		slog.Error("Could not start server")
	}
}

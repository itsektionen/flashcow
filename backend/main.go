package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/itsektionen/flashcow/backend/internal/api"
	"github.com/itsektionen/flashcow/backend/internal/logic"
	"github.com/itsektionen/flashcow/backend/internal/storage"

	"github.com/jackc/pgx/v5"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	defer conn.Close(context.Background())

	if err != nil {
		slog.Error("Could not connect to database")
		os.Exit(1)
	}

	r := &storage.Database{
		Conn: conn,
	}

	s := &logic.Service{
		Repo: r,
	}

	srv := &api.Server{
		Addr:    ":80",
		Service: s,
	}

	if err := srv.Serve(); err != nil {
		slog.Error("Could not start server")
	}
}

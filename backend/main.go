package main

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/itsektionen/flashcow/backend/internal/api"
	"github.com/itsektionen/flashcow/backend/internal/repository"
	"github.com/itsektionen/flashcow/backend/internal/service"

	// import the dialect in both goqu and sql (pq)
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	//dbConn, err := sql.Open("postgres", os.Getenv("POSTGRES_DSN"))
	dbConn, err := sql.Open("postgres", "postgres://dev:dev@127.0.0.1:5432/flashcow?sslmode=disable")
	if err != nil {
		slog.Error("Could not connect to database", "error", err)
		os.Exit(1)
	}

	defer func() {
		_ = dbConn.Close()
	}()

	db := &repository.Database{
		Conn: dbConn,
	}

	err = db.Migrate()
	if err != nil {
		slog.Error("Could not migrate database", "error", err)
		os.Exit(1)
	}

	userRepository := &repository.UserRepository{
		Db: db,
	}

	committeeRepository := &repository.CommitteeRepository{}

	userService := &service.UserService{
		Repository: userRepository,
	}

	committeeService := &service.CommitteeService{
		Repository: committeeRepository,
	}

	srv := &api.Server{
		Addr:             ":8080",
		UserService:      userService,
		CommitteeService: committeeService,
	}

	if err := srv.Serve(); err != nil {
		slog.Error("Could not start server")
		os.Exit(1)
	}
}

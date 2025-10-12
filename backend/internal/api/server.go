package api

import (
	"log/slog"
	"net/http"

	"github.com/itsektionen/flashcow/backend/internal/logic"
)

type Server struct {
	Addr    string
	Service *logic.Service
}

func (s *Server) Serve() error {
	srv := http.Server{
		Addr:    s.Addr,
		Handler: s.Handler(),
	}

	slog.Info("Starting listener", "addr", s.Addr)
	return srv.ListenAndServe()
}

func (s Server) Handler() http.Handler {
	mux := http.NewServeMux()

	userHandler := UserHandler{Service: s.Service}
	committeeHandler := CommitteeHandler{}

	mux.HandleFunc("GET /api/user", userHandler.listUsers)
	mux.HandleFunc("POST /api/user", userHandler.createUser)

	mux.HandleFunc("GET /api/user/{id}", userHandler.getUser)
	mux.HandleFunc("PUT /api/user/{id}", userHandler.updateUser)
	mux.HandleFunc("DELETE /api/user/{id}", userHandler.deleteUser)

	mux.HandleFunc("GET /api/committee", committeeHandler.getCommittees)

	return mux
}

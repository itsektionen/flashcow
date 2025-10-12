package logic

import "github.com/itsektionen/flashcow/backend/internal/storage"

type Service struct {
	Repo *storage.Database
}

package logic

import (
	"github.com/itsektionen/flashcow/backend/internal/repository"
)

type Service struct {
	Repo *repository.Database
}

package repository

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/itsektionen/flashcow/backend/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id int64) (*model.User, error)
	ListUsers(ctx context.Context) ([]model.User, error)
}

type userRepository struct {
	Db *Database
}

func NewUserRepository(db *Database) UserRepository {
	return &userRepository{
		Db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	exec := r.Db.handle().Insert("user_details").Cols("full_name", "chapter_email").Vals(
		goqu.Vals{user.FullName, user.ChapterEmailAddress}).Returning(
		goqu.L("id")).Executor()

	var id int64
	if found, err := exec.ScanValContext(ctx, &id); err != nil {
		return err
	} else if found == false {
		return ErrNotFound
	}

	user.ID = id

	return nil
}

func (r *userRepository) GetUser(ctx context.Context, id int64) (*model.User, error) {
	exec := r.Db.handle().From("user_details").Select("id", "full_name", "chapter_email").Where(goqu.Ex{"id": id}).Executor()

	user := &model.User{}
	if found, err := exec.ScanStructContext(ctx, user); err != nil {
		return nil, err
	} else if found == false {
		return nil, ErrNotFound
	}

	return user, nil
}

func (r *userRepository) ListUsers(ctx context.Context) ([]model.User, error) {
	exec := r.Db.handle().From("user_details").Select("id", "full_name", "chapter_email").Executor()

	users := []model.User{}
	if err := exec.ScanStructsContext(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

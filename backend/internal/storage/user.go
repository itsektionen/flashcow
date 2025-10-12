package storage

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/itsektionen/flashcow/backend/internal"
)

func (db *Database) CreateUser(ctx context.Context, user *internal.User) error {
	exec := db.handle().Insert("user_details").Cols("full_name", "chapter_email").Vals(
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

func (db *Database) GetUser(ctx context.Context, id int64) (*internal.User, error) {
	exec := db.handle().From("user_details").Select("id", "full_name", "chapter_email").Where(goqu.Ex{"id": id}).Executor()

	user := &internal.User{}
	if found, err := exec.ScanStructContext(ctx, user); err != nil {
		return nil, err
	} else if found == false {
		return nil, ErrNotFound
	}

	return user, nil
}

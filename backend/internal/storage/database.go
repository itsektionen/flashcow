package storage

import (
	"context"
	"log/slog"

	"github.com/doug-martin/goqu/v9"
	"github.com/itsektionen/flashcow/backend/internal"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	Conn *pgx.Conn
}

func (db *Database) CreateUser(user *internal.User) error {
	ds := goqu.Insert("user_details").Cols("full_name", "chapter_email").Vals(
		goqu.Vals{user.FullName, user.ChapterEmailAddress}).Returning(
		goqu.L("id"))

	query, args, err := ds.ToSQL()

	if err != nil {
		return err
	}

	slog.Info("Executing query", "query", query, "args", args)

	var id int64
	err = db.Conn.QueryRow(context.TODO(), query, args...).Scan(&id)
	if err != nil {
		slog.Error("User couldn't be created in the database")
		return err
	}

	user.ID = id

	return nil
}

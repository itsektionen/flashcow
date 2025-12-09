package repository

import (
	"database/sql"
	"embed"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"sort"

	"github.com/facette/natsort"
	"github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

//go:embed migrations/*.sql
var fs embed.FS

type natSort []os.DirEntry

func (n natSort) Len() int {
	return len(n)
}

func (n natSort) Less(i, j int) bool {
	return natsort.Compare(n[i].Name(), n[j].Name())
}

func (n natSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[j]
}

func (db *Database) Migrate() error {
	files, err := fs.ReadDir("migrations")
	if err != nil {
		return errors.Join(errors.New("could not read directory"), err)
	}

	sort.Sort(natSort(files))

	var m []string

	err = db.handle().From("migrations").Select("filename").Executor().ScanVals(&m)
	if err != nil && !isErrorTableNotExist(err) {
		return err
	}

	for _, f := range files {
		if slices.Contains(m, f.Name()) {
			continue
		}

		if err := db.migrateFile(f.Name()); err != nil {
			return errors.Join(errors.New("could not migrate file"), err)
		}
	}

	slog.Info("successfully migrated database")
	return nil
}

func isErrorTableNotExist(err error) bool {
	var pqErr *pq.Error

	// https://www.postgresql.org/docs/current/errcodes-appendix.html
	if errors.As(err, &pqErr) {
		return pqErr.Code == "42P01"
	}
	return false
}

func (db *Database) migrateFile(name string) error {
	slog.Debug("running migration file", "file", name)

	data, err := fs.ReadFile(filepath.Join("migrations", name))
	if err != nil {
		return err
	}

	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	if _, err = tx.Exec(string(data)); err != nil {
		// defer handles rollback
		return err
	}

	if _, err := tx.Exec("INSERT INTO migrations (filename) VALUES ($1)", name); err != nil {
		// defer handles rollback
		return errors.Join(errors.New("could not insert to migrations"), err)
	}

	if err := tx.Commit(); err != nil {
		return errors.Join(errors.New("could not commit migration"), err)
	}

	return nil
}

package repository

import "github.com/doug-martin/goqu/v9"

func (db *Database) handle() *goqu.Database {
	return goqu.Dialect("postgres").DB(db.Conn)
}

package dao

import (
	"database/sql"
)

type data struct {
	*sql.DB
}

func (u *data) Init(db *sql.DB) (err error) {
	u.DB = db
	return nil
}

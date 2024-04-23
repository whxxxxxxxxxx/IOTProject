package dao

import (
	"database/sql"
)

var (
	Data = &data{}
)

func InitMS(db *sql.DB) error {
	err := Data.Init(db)
	if err != nil {
		return err
	}

	return err
}

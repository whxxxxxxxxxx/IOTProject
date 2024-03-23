package dao

import (
	"gorm.io/gorm"
)

var (
	Device = &device{}
)

func InitMS(db *gorm.DB) error {
	err := Device.Init(db)
	if err != nil {
		return err
	}

	return err
}

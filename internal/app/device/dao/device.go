package dao

import (
	"IOTProject/internal/app/device/model"
	"gorm.io/gorm"
)

type device struct {
	*gorm.DB
}

func (u *device) Init(db *gorm.DB) (err error) {
	u.DB = db
	return db.AutoMigrate(&model.Device{}, &model.Location{}, &model.NetworkInfo{}, &model.Security{})
}

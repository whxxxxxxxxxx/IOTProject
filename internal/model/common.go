package model

import (
	"gorm.io/gorm"
	"time"
)

type GormBase struct {
	ID        string         `json:"id" gorm:"primary_key;type:char(26)"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

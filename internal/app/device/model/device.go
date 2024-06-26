package model

import (
	"IOTProject/internal/model"
	"gorm.io/gorm"
	"time"
)

// 数据库模型

type Device struct {
	model.Base
	DeviceNo     string `gorm:"comment:设备编号"`
	Name         string `gorm:"comment:设备名称"`
	ModelData    string `gorm:"comment:设备型号"`
	SerialNumber string `gorm:"comment:设备序列号"`
	//has one
	Location    Location
	NetworkInfo NetworkInfo
	Security    Security
}

type Location struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	DeviceID  string         `gorm:"primaryKey;comment:设备ID"`
	Latitude  float64        `gorm:"comment:纬度"`
	Longitude float64        `gorm:"comment:经度"`
}

type NetworkInfo struct {
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	DeviceID   string         `gorm:"primaryKey;comment:设备ID"`
	IPAddress  string         `gorm:"comment:IP地址"`
	MacAddress string         `gorm:"comment:MAC地址"`
}

type Security struct {
	CreatedAt            time.Time      `json:"-"`
	UpdatedAt            time.Time      `json:"-"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
	DeviceID             string         `gorm:"primaryKey;comment:设备ID"`
	EncryptionStatus     string         `gorm:"comment:加密状态"`
	AuthenticationMethod string         `gorm:"comment:认证方式"`
}

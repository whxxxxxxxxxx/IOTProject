package dto

import (
	"IOTProject/internal/app/data/model"
	"time"
)

type Data struct {
	DeviceID           string                   `json:"deviceId"`
	Status             model.Status             `json:"status"`
	PerformanceMetrics model.PerformanceMetrics `json:"performanceMetrics"`
	TimeStamp          int64                    `json:"timeStamp"`
}

type PerformanceMetricsString struct {
	CreateTime  time.Time
	Temperature string
	Pressure    string
	Speed       string
	Output      string
}

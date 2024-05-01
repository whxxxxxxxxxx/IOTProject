package model

import "time"

// 数据库模型

type Status struct {
	CreateTime        time.Time
	PowerState        string `sql:"VARCHAR(16)"`
	OperationalStatus string `sql:"VARCHAR(16)"`
	Mode              string `sql:"VARCHAR(16)"`
}

type PerformanceMetrics struct {
	CreateTime  time.Time
	Temperature float64
	Pressure    float64
	Speed       float64
	Output      float64
}

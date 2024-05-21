package model

import "time"

// 数据库模型

type Status struct {
	Ts                time.Time
	PowerState        string `sql:"VARCHAR(16)"`
	OperationalStatus string `sql:"VARCHAR(16)"`
	Mode              string `sql:"VARCHAR(16)"`
}

type PerformanceMetrics struct {
	Ts          time.Time
	Temperature float64
	Pressure    float64
	Speed       float64
	Output      float64
}

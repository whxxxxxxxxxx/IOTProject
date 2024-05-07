package dao

import (
	"IOTProject/internal/app/data/model"
	"IOTProject/internal/app/data/service"
	"database/sql"
	"fmt"
)

type data struct {
	*sql.DB
}

func (u *data) Init(db *sql.DB) (err error) {
	u.DB = db
	err = CreateStable(db, "Status", model.Status{})
	if err != nil {
		return err
	}
	err = CreateStable(db, "PerformanceMetrics", model.PerformanceMetrics{})
	if err != nil {
		return err
	}
	return nil
}

func (u *data) GetDataByIdAndCurrentTime(id string, timeFromNow string) (interface{}, error) {
	interval, err := service.ConvertTimeToSeconds(timeFromNow)
	if err != nil {
		return nil, err
	}
	sqlSentence := fmt.Sprintf("SELECT * FROM %s WHERE ts >=NOW - %s INTERVAL(%s)", id, timeFromNow, interval)
	if id[0] == 's' {
		var statusDatas []model.Status
		//SELECT * FROM tb1 WHERE ts >= NOW - 1h;
		rows, err := u.DB.Query(sqlSentence)
		//select _wstart, tbname, avg(voltage) from meters partition by tbname interval(10m)
		for rows.Next() {
			var status model.Status
			err = rows.Scan(&status.CreateTime, &status.PowerState, &status.OperationalStatus, &status.Mode)
			if err != nil {
				return nil, err
			}
			statusDatas = append(statusDatas, status)
		}
		return statusDatas, err
	} else if id[0] == 'p' {
		var performanceMetricsDatas []model.PerformanceMetrics
		rows, err := u.DB.Query(sqlSentence)
		for rows.Next() {
			var performanceMetrics model.PerformanceMetrics
			err = rows.Scan(&performanceMetrics.CreateTime, &performanceMetrics.CreateTime, &performanceMetrics.Pressure, &performanceMetrics.Output, &performanceMetrics.Speed, &performanceMetrics.Temperature)
			if err != nil {
				return nil, err
			}
			performanceMetricsDatas = append(performanceMetricsDatas, performanceMetrics)
		}
		return performanceMetricsDatas, err
	}
	return nil, nil
}

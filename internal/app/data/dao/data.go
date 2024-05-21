package dao

import (
	"IOTProject/internal/app/data/dto"
	"IOTProject/internal/app/data/model"
	"IOTProject/pkg/timex"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
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
	interval, err := timex.ConvertTimeToSeconds(timeFromNow)
	if err != nil {
		return nil, err
	}

	if id[0] == 's' {
		selectContent := "MODE(PowerState),MODE(OperationalStatus),MODE(Mode)"
		sqlSentence := fmt.Sprintf("SELECT _wstart,%s FROM %s WHERE ts >=NOW - %s INTERVAL(%s)", selectContent, id, timeFromNow, interval)
		var statusDatas []model.Status
		//SELECT * FROM tb1 WHERE ts >= NOW - 1h;
		rows, err := u.DB.Query(sqlSentence)
		//select _wstart, tbname, avg(voltage) from meters partition by tbname interval(10m)
		for rows.Next() {
			var status model.Status
			err = rows.Scan(&status.Ts, &status.PowerState, &status.OperationalStatus, &status.Mode)
			if err != nil {
				return nil, err
			}
			statusDatas = append(statusDatas, status)
		}
		return statusDatas, err
	} else if id[0] == 'p' {
		selectContent := "AVG(Temperature),AVG(Pressure),AVG(Speed),AVG(Output)"
		sqlSentence := fmt.Sprintf("SELECT _wstart,%s FROM %s WHERE ts >=NOW - %s INTERVAL(%s)", selectContent, id, timeFromNow, interval)
		var performanceMetricsDatas []model.PerformanceMetrics
		rows, err := u.DB.Query(sqlSentence)
		for rows.Next() {
			var performanceMetrics model.PerformanceMetrics
			err = rows.Scan(&performanceMetrics.Ts, &performanceMetrics.Pressure, &performanceMetrics.Output, &performanceMetrics.Speed, &performanceMetrics.Temperature)
			if err != nil {
				return nil, err
			}
			performanceMetricsDatas = append(performanceMetricsDatas, performanceMetrics)
		}
		return performanceMetricsDatas, err
	}
	return nil, nil
}

func (u *data) InsertData(id string, data interface{}) error {
	val := reflect.ValueOf(data)
	stableName := val.Type().Name()
	fieldNum := val.NumField()
	values := make([]string, fieldNum)
	for i := 0; i < fieldNum; i++ {
		fieldValue := val.Field(i)
		var valueStr string
		switch fieldValue.Kind() {
		case reflect.String:
			valueStr = fmt.Sprintf("'%s'", fieldValue.String())
		case reflect.Float64:
			valueStr = fmt.Sprintf("%f", fieldValue.Float())
		case reflect.Int:
			valueStr = fmt.Sprintf("%d", fieldValue.Int())
		case reflect.Struct:
			if fieldValue.Type().String() == "time.Time" {
				// 格式化时间为SQL兼容格式
				valueStr = fmt.Sprintf("'%s'", fieldValue.Interface().(time.Time).Format("2006-01-02 15:04:05"))
			}
		default:
			valueStr = fmt.Sprintf("'%s'", fieldValue.String())
		}
		values[i] = valueStr
	}
	sqlSentence := fmt.Sprintf("INSERT INTO %s USING %s TAGS('China',1) VALUES (%s)",
		id, stableName, strings.Join(values, ", "))
	_, err := u.DB.Exec(sqlSentence)
	return err
}

func (u *data) InsertStatusData(id string, status model.Status) error {
	sqlSentence := fmt.Sprintf("INSERT INTO %s (Ts, PowerState, OperationalStatus, Mode) USING status TAGS('China',1) VALUES (?, ?, ?, ?) ", id)
	_, err := u.DB.Exec(sqlSentence, status.Ts, status.PowerState, status.OperationalStatus, status.Mode)
	return err
}

func (u *data) InsertPerformanceMetricsData(id string, performanceMetrics model.PerformanceMetrics) error {
	sqlSentence := fmt.Sprintf("INSERT INTO %s (Ts, Temperature, Pressure, Speed, Output) USING performancemetrics TAGS('China',1) VALUES (?, ?, ?, ?, ?) ", id)
	_, err := u.DB.Exec(sqlSentence, performanceMetrics.Ts, performanceMetrics.Temperature, performanceMetrics.Pressure, performanceMetrics.Speed, performanceMetrics.Output)
	return err
}

func (u *data) InsertDataById(Data dto.Data) error {
	deviceId := Data.DeviceID
	statusId := "s" + deviceId
	performanceMetricsId := "p" + deviceId
	//插入状态数据
	err := u.InsertData(statusId, Data.Status)
	if err != nil {
		return err
	}
	//插入性能数据
	err = u.InsertData(performanceMetricsId, Data.PerformanceMetrics)
	if err != nil {
		return err
	}
	return nil
}

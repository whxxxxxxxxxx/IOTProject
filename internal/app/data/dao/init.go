package dao

import (
	"database/sql"
	"fmt"
	"reflect"
)

var (
	Data = &data{}
)

func InitTD(db *sql.DB) error {
	err := Data.Init(db)
	if err != nil {
		return err
	}

	return err
}

func CreateStable(db *sql.DB, tableName string, data interface{}) error {
	sqlSentence := GenerateCreateTableSQL(tableName, data)
	_, err := db.Exec(sqlSentence)
	return err
}

func GenerateCreateTableSQL(tableName string, data interface{}) string {
	sqlSentence := fmt.Sprintf("CREATE STABLE IF NOT EXISTS %s (ts TIMESTAMP,", tableName)
	dataType := reflect.TypeOf(data)
	for i := 0; i < dataType.NumField(); i++ {
		fieldName := dataType.Field(i).Name
		fieldType := dataType.Field(i).Type
		sqlTag := dataType.Field(i).Tag.Get("sql")
		sqlType := ""
		if sqlTag == "" {
			switch fieldType.Kind() {
			case reflect.String:
				sqlType = fmt.Sprintf("BINARY(%d)", 8)
			case reflect.Float64:
				sqlType = "FLOAT"
			case reflect.Int:
				sqlType = "INT"
			default:
				continue
			}
		} else {
			sqlType = sqlTag
		}
		sqlSentence += fmt.Sprintf("%s %s,", fieldName, sqlType)
	}

	sqlSentence = sqlSentence[:len(sqlSentence)-1] + ")" + " TAGS (location BINARY(64), groupId INT)"
	return sqlSentence
}

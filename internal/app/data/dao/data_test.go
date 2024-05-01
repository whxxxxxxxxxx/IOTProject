package dao

import (
	"reflect"
	"testing"
)

type test struct {
	Name  string `sql:"VARCHAR(16)"`
	Age   int
	Score float64
}

func TestGenerateCreateTableSQL(t *testing.T) {
	var test test
	sql := GenerateCreateTableSQL("test", test)
	t.Log(sql)
}

func TestTag(t *testing.T) {
	dataType := reflect.TypeOf(test{})
	for i := 0; i < dataType.NumField(); i++ {
		fieldName := dataType.Field(i).Name
		//fieldType := dataType.Field(i).Type
		sqlTag := dataType.Field(i).Tag.Get("sql")
		if sqlTag == "" {
			print(fieldName, " is not a column")
		} else {
			print(fieldName, sqlTag)
		}
	}
}

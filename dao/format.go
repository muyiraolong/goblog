package dao

import (
	"database/sql"
	"reflect"
	"strconv"
	"time"
)

type ForMatDB struct {
	*sql.DB
}

var MatDB ForMatDB

func (d *ForMatDB) QueryOne(model interface{}, sql string, args ...interface{}) error {
	rows, err := d.Query(sql, args...)
	if err != nil {
		return err
	}
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	vals := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	for k := range vals {
		scans[k] = &vals[k]
	}
	if rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return err
		}
	}
	var result = make(map[string]interface{})
	elem := reflect.ValueOf(model).Elem()
	for index, val := range columns {
		result[val] = string(vals[index])
	}
	for i := 0; i < elem.NumField(); i++ {
		structField := elem.Type().Field(i)
		fieldInfo := structField.Tag.Get("orm")
		v := result[fieldInfo]
		t := structField.Type
		switch t.String() {
		case "int":
			s := v.(string)
			vInt, _ := strconv.Atoi(s)
			elem.Field(i).Set(reflect.ValueOf(vInt))
		case "string":
			elem.Field(i).Set(reflect.ValueOf(v.(string)))
		case "int64":
			s := v.(string)
			vInt64, _ := strconv.ParseInt(s, 10, 64)
			elem.Field(i).Set(reflect.ValueOf(vInt64))
		case "int32":
			s := v.(string)
			vInt32, _ := strconv.ParseInt(s, 10, 32)
			elem.Field(i).Set(reflect.ValueOf(vInt32))
		case "time.Time":
			s := v.(string)
			t, _ := time.Parse(time.RFC3339, s)
			elem.Field(i).Set(reflect.ValueOf(t))
		}
	}
	return nil
}

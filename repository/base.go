package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	// "radidev/database"
)

type Anda struct {
	Id           string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

func ConvertStructItemToMap[T any](f T) map[string]any {
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	if v.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return nil
	}

	data := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		data[fieldName] = fieldValue
	}
	return data

}

func Create(db *sql.DB, tableName string, data map[string]any) (string, error) {
	cols := []string{}
	vals := []any{}
	placeholders := []string{}

	for col, val := range data {
		cols = append(cols, col)
		vals = append(vals, val)
	}

	for i := 1; i <= len(cols); i++ {
		placeholders = append(placeholders, "$"+strconv.Itoa(i))
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", tableName, strings.Join(cols, ", "),
		strings.Join(placeholders, ", "))

	var id string
	err := db.QueryRow(query, vals...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func ListAsMaps(db *sql.DB, tableName string, fields ...string) ([]map[string]any, error) {
	columns := "*"
	if len(fields) > 0 {
		columns = strings.Join(fields, ", ")
	}
	query := fmt.Sprintf("SELECT %s FROM %s", columns, tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columnNames, _ := rows.Columns()
	values := make([]any, len(columnNames))
	valuePtrs := make([]any, len(columnNames))

	var result []map[string]any

	for rows.Next() {
		for i := range columnNames {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]any)
		for i, col := range columnNames {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}
		result = append(result, rowMap)
	}

	return result, nil
}

// Fields of T have to match fields in size and order
func GetById[T any](db *sql.DB, tableStruct T, tableName string, id string, fields ...string) (T, error) {
	columns := "*"
	if len(fields) > 0 {
		columns = strings.Join(fields, ", ")
	}
	query := fmt.Sprintf("SELECT %s from %s WHERE id=($1)", columns, tableName)
	var item T
	t := reflect.TypeOf(item)
	tPtr := reflect.New(t)
	v := tPtr.Elem()
	if len(fields) >= 1 && v.NumField() != len(fields) {
		return item, fmt.Errorf("length of %s: %d, is not the same with length of fields provided:%d", t, v.NumField(), len(fields))
	}

	var fieldPtrs []any
	for i := 0; i < v.NumField(); i++ {
		fieldPtrs = append(fieldPtrs, v.Field(i).Addr().Interface())
	}

	err := db.QueryRow(query, id).Scan(fieldPtrs...)

	if err != nil {
		return item, err
	}
	return tPtr.Elem().Interface().(T), nil
}

func Delete(db *sql.DB, tableName string, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=($1)", tableName)
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

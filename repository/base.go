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
		fmt.Printf("%s: %v...\n", fieldName, fieldValue)
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

	fmt.Println("Executing query:", query)
	var id string
	err := db.QueryRow(query, vals...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// func List[T any](db *sql.DB, table string, fields ...string) ([]T, error) {
// 	columns := "*"
// 	if len(fields) > 0 {
// 		columns = strings.Join(fields, ", ")
// 	}
// 	query := fmt.Sprintf("SELECT %s from %s", columns, table)
// 	fmt.Println("Executing query:", query)
// 	rows, err := db.Query(query)
// 	fmt.Println("Rows:", rows)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	tType := reflect.TypeOf((*T)(nil)).Elem()
// 	var results []T

// 	for rows.Next() {
// 		tPtr := reflect.New(tType)
// 		tVal := tPtr.Elem()
// 		var fieldPtrs []any
// 		for i := 0; i < tVal.NumField(); i++ {
// 			fieldPtrs = append(fieldPtrs, tVal.Field(i).Addr().Interface())
// 		}

// 		// for i := 0; i < len(fields); i++ {
// 		// 	fieldPointers = append(fieldPointers, reflect.New(reflect.TypeOf(u).Field(i).Type).Interface())
// 		fmt.Println("Field pointers:", fieldPtrs)
// 		err := rows.Scan(fieldPtrs...)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Convert pointer to value
// 		results = append(results, tPtr.Interface().(T))
// 	}

// 	return results, rows.Err()
// }

func ListAsMaps(db *sql.DB, table string, fields ...string) ([]map[string]any, error) {
	columns := "*"
	if len(fields) > 0 {
		columns = strings.Join(fields, ", ")
	}
	query := fmt.Sprintf("SELECT %s FROM %s", columns, table)
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

// func (user User) Get(db *sql.DB, id string) (User, error) {
// 	query := `SELECT id, username, password_hash, created_at from users WHERE id=($1)`
// 	var u User
// 	err := db.QueryRow(query, id).Scan(&u.Id, &u.Username, &u.PasswordHash, &u.CreatedAt)
// 	if err != nil {
// 		return u, err
// 	}
// 	return u, nil
// }

// func (user User) Delete(db *sql.DB, id string) error {
// 	query := `DELETE FROM users WHERE id=($1)`
// 	_, err := db.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

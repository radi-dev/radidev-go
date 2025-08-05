package repository

import (
	"database/sql"
	"time"
	// "radidev/database"
)

type User struct {
	Id           string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

var tableNameUser = "users"

func (user *User) CreateUser(db *sql.DB, data map[string]any) (string, error) {
	return Create(db, tableNameUser, data)
}

func (user User) ListUsers(db *sql.DB, fields ...string) ([]map[string]any, error) {
	return ListAsMaps(db, tableNameUser, fields...)

}

func (user User) GetUser(db *sql.DB, userStruct User, id string, fields ...string) (User, error) {
	return GetById(db, userStruct, tableNameUser, id, fields...)
}

func (user User) DeleteUser(db *sql.DB, id string) error {

	return Delete(db, tableNameUser, id)
}

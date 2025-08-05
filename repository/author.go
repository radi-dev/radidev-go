package repository

import (
	"database/sql"
	"time"
	// "radidev/database"
)

type Author struct {
	Id           string
	Authorname   string
	PasswordHash string
	CreatedAt    time.Time
}

var tableNameAuthor = "authors"

func (author *Author) CreateAuthor(db *sql.DB, data map[string]any) (string, error) {
	return Create(db, tableNameAuthor, data)
}

func (author Author) ListAuthors(db *sql.DB, fields ...string) ([]map[string]any, error) {
	return ListAsMaps(db, tableNameAuthor, fields...)

}

func (author Author) GetAuthor(db *sql.DB, authorStruct Author, id string, fields ...string) (Author, error) {
	return GetById(db, authorStruct, tableNameAuthor, id, fields...)
}

func (author Author) DeleteAuthor(db *sql.DB, id string) error {

	return Delete(db, tableNameAuthor, id)
}

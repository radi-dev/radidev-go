package repository

import (
	"database/sql"
	"time"
	// "radidev/database"
)

type Comment struct {
	Id           string
	Commentname  string
	PasswordHash string
	CreatedAt    time.Time
}

var tableNameComment = "comments"

func (comment *Comment) CreateComment(db *sql.DB, data map[string]any) (string, error) {
	return Create(db, tableNameComment, data)
}

func (comment Comment) ListComments(db *sql.DB, fields ...string) ([]map[string]any, error) {
	return ListAsMaps(db, tableNameComment, fields...)

}

func (comment Comment) GetComment(db *sql.DB, commentStruct Comment, id string, fields ...string) (Comment, error) {
	return GetById(db, commentStruct, tableNameComment, id, fields...)
}

func (comment Comment) DeleteComment(db *sql.DB, id string) error {

	return Delete(db, tableNameComment, id)
}

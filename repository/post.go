package repository

import (
	"database/sql"
	"time"
	// "radidev/database"
)

type Post struct {
	Id           string
	Postname     string
	PasswordHash string
	CreatedAt    time.Time
}

var tableNamePost = "posts"

func (post *Post) CreatePost(db *sql.DB, data map[string]any) (string, error) {
	return Create(db, tableNamePost, data)
}

func (post Post) ListPosts(db *sql.DB, fields ...string) ([]map[string]any, error) {
	return ListAsMaps(db, tableNamePost, fields...)

}

func (post Post) GetPost(db *sql.DB, postStruct Post, id string, fields ...string) (Post, error) {
	return GetById(db, postStruct, tableNamePost, id, fields...)
}

func (post Post) DeletePost(db *sql.DB, id string) error {

	return Delete(db, tableNamePost, id)
}

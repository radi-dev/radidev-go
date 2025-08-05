package repository

import (
	"database/sql"
	"time"
	// "radidev/database"
)

type Tag struct {
	Id           string
	Tagname      string
	PasswordHash string
	CreatedAt    time.Time
}

var tableNameTag = "tags"

func (tag *Tag) CreateTag(db *sql.DB, data map[string]any) (string, error) {
	return Create(db, tableNameTag, data)
}

func (tag Tag) ListTags(db *sql.DB, fields ...string) ([]map[string]any, error) {
	return ListAsMaps(db, tableNameTag, fields...)

}

func (tag Tag) GetTag(db *sql.DB, tagStruct Tag, id string, fields ...string) (Tag, error) {
	return GetById(db, tagStruct, tableNameTag, id, fields...)
}

func (tag Tag) DeleteTag(db *sql.DB, id string) error {

	return Delete(db, tableNameTag, id)
}

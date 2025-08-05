package repository

import (
	"database/sql"
	"time"
	// "radidev/database"
)

type Upload struct {
	Id           string
	Uploadname   string
	PasswordHash string
	CreatedAt    time.Time
}

var tableNameUpload = "uploads"

func (upload *Upload) CreateUpload(db *sql.DB, data map[string]any) (string, error) {
	return Create(db, tableNameUpload, data)
}

func (upload Upload) ListUploads(db *sql.DB, fields ...string) ([]map[string]any, error) {
	return ListAsMaps(db, tableNameUpload, fields...)

}

func (upload Upload) GetUpload(db *sql.DB, uploadStruct Upload, id string, fields ...string) (Upload, error) {
	return GetById(db, uploadStruct, tableNameUpload, id, fields...)
}

func (upload Upload) DeleteUpload(db *sql.DB, id string) error {

	return Delete(db, tableNameUpload, id)
}

package database

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"

	"radidev/config"
)

var cfg = config.Load()

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}
	log.Printf("Connected to the database successfully: %d", db.Stats().OpenConnections)
	return db, nil
}

func CreateTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS authors (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150) UNIQUE NOT NULL,
		about TEXT
	);


	CREATE TABLE IF NOT EXISTS posts (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(255) NOT NULL,
		is_project BOOLEAN DEFAULT FALSE,
		is_published BOOLEAN DEFAULT FALSE,
		tags JSONB DEFAULT '[]',
		author_id UUID REFERENCES authors(id) ON DELETE SET NULL,
		text TEXT NOT NULL,
		featured_image_url TEXT,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW()
	);


	CREATE TABLE IF NOT EXISTS comments (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		text TEXT NOT NULL,
		post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
		comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
		author_name VARCHAR(100) NOT NULL,
		author_email VARCHAR(150),
		created_at TIMESTAMPTZ DEFAULT NOW(),
		CHECK (
			(post_id IS NOT NULL AND comment_id IS NULL) OR
			(post_id IS NULL AND comment_id IS NOT NULL)
		)
	);

	CREATE TABLE IF NOT EXISTS uploads (
		upload_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255),
		url TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS tags (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(100) UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS post_tags (
		post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
		tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
		PRIMARY KEY (post_id, tag_id)
	);

	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}
	return nil
}

package repository

import (
	"database/sql"
	"fmt"
	"time"
	// "radidev/database"
)

type User struct {
	Id           string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
}

func (user *User) Create(db *sql.DB) (string, error) {
	query := `INSERT INTO users (username, password_hash, created_at) VALUES ($1, $2, NOW()) RETURNING id`
	var id string
	err := db.QueryRow(query, user.Username, user.PasswordHash).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (user User) List(db *sql.DB) ([]User, error) {
	query := `SELECT id, username, created_at from users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Username, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		fmt.Println("User is", u)
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (user User) Get(db *sql.DB, id string) (User, error) {
	query := `SELECT id, username, password_hash, created_at from users WHERE id=($1)`
	var u User
	err := db.QueryRow(query, id).Scan(&u.Id, &u.Username, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return u, err
	}
	return u, nil
}

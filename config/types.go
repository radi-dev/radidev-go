package config

import "database/sql"

type App struct {
	DB *sql.DB
}

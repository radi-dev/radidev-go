package config

import (
	"log"
	"os"
	"path/filepath"

	"radidev/util"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() Config {
	execDir := util.GetExecutableDir()
	envPath := filepath.Join(execDir, ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("No .env file found or failed to load it", err)
	}

	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}

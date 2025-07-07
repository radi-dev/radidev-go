package config

import (
	"log"
	"os"

	"radidev/util"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		err = godotenv.Load(util.GetAbsPath(".env"))
		if err != nil {
			log.Println("No .env file found or failed to load it", err)
		}
	}

	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}

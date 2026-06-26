package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT       string
	DSN        string
	JWT_SECRET string
}

func LoadEnv() *Env {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env file")
	}

	return &Env{
		PORT:       os.Getenv("PORT"),
		DSN:        os.Getenv("DSN"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
}

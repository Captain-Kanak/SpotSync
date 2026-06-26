package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(env *Env) *gorm.DB {
	db, err := gorm.Open(postgres.Open(env.DSN), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("Database connection failed")
	}

	fmt.Println("Database connection established")

	return db
}

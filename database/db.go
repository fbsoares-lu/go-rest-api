package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	con, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Connection error")
	}

	DB = con
}

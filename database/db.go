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

func ConnectionDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=cbgomes-personas port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Error conneting database...", err)
	}
}

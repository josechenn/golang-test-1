package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	godotenv.Load()

	db, err := sql.Open("mysql", os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GormConnect() *gorm.DB {

	var err error
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)
	return db
}

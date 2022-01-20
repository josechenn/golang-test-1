package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test-golang-1")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GormConnect() *gorm.DB {

	var err error
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test-golang-1")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)
	return db
}

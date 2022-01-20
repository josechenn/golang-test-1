package config

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const projectDirName = "golang-test-1"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func GormConnect() *gorm.DB {
	loadEnv()
	dbUrl := os.Getenv("DB_URL")

	var err error
	db, err := gorm.Open("mysql", dbUrl)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return db
}

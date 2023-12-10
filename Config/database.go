package Config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnection() {
	var err error
	dbUrl := os.Getenv("DB_URL")                              // Getting DB URL from environment variable
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{}) // Open database connection
	if err != nil {
		log.Fatal("Failed to connect to database")
	} else {
		fmt.Println("connecting to database")
	}
}

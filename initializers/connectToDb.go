package initializers

import (
	"log"
	"os"

	"github.com/chtiwa/go_pagination/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	DB.AutoMigrate(&models.Person{})
}

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed database connection!")
	}
}

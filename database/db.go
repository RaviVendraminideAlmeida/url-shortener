package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/RaviVendraminideAlmeida/url-shortener/models"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func ConnectDB() {

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Connection to DB failed: ", err, "\n")
	}

	log.Println("Connection to DB established successfully\n")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.URL{})

	DB = DBInstance{
		Db: db,
	}
}

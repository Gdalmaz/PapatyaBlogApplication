package database

import (
	"fmt"
	"log"
	"os"
	"product/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("POSTGRES_AUTH_HOST"), os.Getenv("POSTGRES_AUTH_USER"),
		os.Getenv("POSTGRES_AUTH_PASSWORD"), os.Getenv("POSTGRES_AUTH_DB"), os.Getenv("POSTGRES_AUTH_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("connected success")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	err = db.AutoMigrate(&models.Product{}, &models.ExtraImage{})
	if err != nil {
		log.Fatal("error to migrate step")
	}

	DB = DBInstance{
		Db: db,
	}

}

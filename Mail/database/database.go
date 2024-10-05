package database

import (
	"fmt"
	"log"
	"mail/models"
	"os"

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
	// .env dosyasını yükleyin
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// PostgreSQL bağlantı bilgilerini alın
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

	// Veritabanına bağlan
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database, got error %v", err)
		return
	}

	log.Println("Connected successfully to the database")

	// Logger ayarlarını yap
	db.Logger = logger.Default.LogMode(logger.Info)

	// Otomatik migrasyonları çalıştır
	log.Println("Running migrations")
	err = db.AutoMigrate(&models.Mail{})
	if err != nil {
		log.Fatal("Migration step failed:", err)
	}

	DB = DBInstance{
		Db: db,
	}
}

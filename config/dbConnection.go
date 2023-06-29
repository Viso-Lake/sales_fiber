package config

import (
	"fmt"
	"os"

	"sales_fiber/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	godotenv.Load()
	dbname := os.Getenv("PG_DB")
	dbuser := os.Getenv("PG_USER")
	dbpass := os.Getenv("PG_PASS")
	dbhost := os.Getenv("PG_HOST")
	dbport := os.Getenv("PG_PORT")

	conn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe",
		dbhost, dbuser, dbpass, dbname, dbport,
	)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	DB = db
	fmt.Printf("DB connection successfully")

	// AutoMigrate(db)
}

func AutoMigrate(conn *gorm.DB) {
	conn.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	)
}

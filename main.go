package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sales_fiber/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func main() {
	godotenv.Load()

	dbname := os.Getenv("PG_DB")
	dbuser := os.Getenv("PG_USER")
	dbpass := os.Getenv("PG_PASS")
	dbhost := os.Getenv("PG_HOST")
	dbport := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe",
		dbhost, dbuser, dbpass, dbname, dbport,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&model.Cachier{})

	DB = db

	fmt.Printf("DB connection successfully")

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Go fiber api created",
		})
	})

	app.Listen(":3000")
}

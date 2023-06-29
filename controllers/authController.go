package controllers

import (
	"os"
	db "sales_fiber/config"
	"sales_fiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Name      string
	Issuer    uint
	ExpiresAt int64
}

func Login(c *fiber.Ctx) error {
	godotenv.Load()

	cashierId := c.Params("cashierId")
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid post request",
		})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Passcode is required",
			"error":   map[string]interface{}{},
		})
	}

	var cashier models.Cashier
	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	if cashier.Passcode != data["passcode"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Passcode not match",
			"error":   map[string]interface{}{},
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, UserClaim{
		Name:      cashier.Name,
		Issuer:    cashier.Id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Token Expired or invalided",
		})
	}

	cashierData := make(map[string]interface{})
	cashierData["token"] = jwtToken

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashierData,
	})
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func Passcode(c *fiber.Ctx) error {
	return nil
}

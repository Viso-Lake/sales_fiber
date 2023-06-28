package controllers

import (
	db "sales_fiber/config"
	"sales_fiber/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier name is required",
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier passcode is required",
			})
	}

	// now saving cashier to db
	cashier := &models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier added successful",
		"data":    cashier,
	})
}

func EditCashier(c *fiber.Ctx) error {
	return nil
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func CashiersList(c *fiber.Ctx) error {
	var count int64
	var cashier []models.Cashier

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier)
	db.DB.Model(&models.Cashier{}).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier list api",
		"data":    cashier,
		"count":   count,
	})
}

func GetCashierDetails(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}

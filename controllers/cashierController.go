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
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
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
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id=?", cashierId)

	// validation for cheching cashier id

	if cashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	var updateCashier models.Cashier
	err := c.BodyParser(&updateCashier)
	if err != nil {
		return err
	}

	if updateCashier.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}

	cashier.Name = updateCashier.Name

	db.DB.Save(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier updated successful",
		"data":    cashier,
	})
}

func CashiersList(c *fiber.Ctx) error {
	var count int64
	var cashier []models.Cashier

	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 20
	}
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

	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Where("id=?", cashierId).First(&cashier)
	// SELECT * FROM cashier WHERE id = cashierId and LIMIT 1 etc
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	db.DB.Where("id=?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier deleted successful",
	})
}

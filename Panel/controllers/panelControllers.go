package controllers

import (
	"panel/database"
	"panel/models"

	"github.com/gofiber/fiber/v2"
)

func GetDeletedProducts(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"status": "error", "Message": "ERROR : G-D-P-1"})
	}
	var products []models.Product
	err := database.DB.Db.Where("is_active = ?", false).Order("Id DESC").Find(&products).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : G-A-P-1"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": products})
}

func GetAvaibleProducts(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"status": "error", "Message": "ERROR : G-D-P-1"})
	}
	var products []models.Product
	err := database.DB.Db.Where("is_active = ?", true).Order("Id DESC").Find(&products).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : G-A-P-1"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": products})
}

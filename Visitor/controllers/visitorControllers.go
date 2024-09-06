package controllers

import (
	"visitor/database"
	"visitor/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllPost(c *fiber.Ctx) error {
	var products []models.Product
	err := database.DB.Db.Where("is_active = ?", true).Order("Id DESC").Find(&products).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : G-A-P-1"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": products})
}

func GetAllExtraImage(c *fiber.Ctx) error {
	var extraimage []models.ExtraImage
	err := database.DB.Db.Where("is_active =?", true).Find(&extraimage).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "Error : G-A-E-I-1"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": extraimage})
}

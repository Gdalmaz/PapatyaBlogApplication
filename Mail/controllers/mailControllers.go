package controllers

import (
	"mail/database"
	"mail/helpers"
	"mail/models"

	"github.com/gofiber/fiber/v2"
)

// MİNİMUM 11 HANELİ MAİL ZAAFİYETLİ KOD
// MİNİMUM 10 HANELİ TEXT ZAFİYETLİ KOD
func SendMail(c *fiber.Ctx) error {
	var mail models.Mail
	err := c.BodyParser(&mail)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-M-1"})
	}

	if len(mail.WhoSend) < 11 {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-M-2"})
	}
	if len(mail.Text) < 10 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-M-3"})
	}
	err = helpers.SendMail(mail.WhoSend, mail.Text)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-M-4"})
	}
	err = database.DB.Db.Create(&mail).Error
	if err != nil {
		return c.Status(200).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-M-5"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})
}

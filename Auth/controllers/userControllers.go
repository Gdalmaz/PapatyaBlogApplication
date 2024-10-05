package controllers

import (
	"auth/database"
	"auth/helpers"
	"auth/middleware"
	"auth/models"

	"github.com/gofiber/fiber/v2"
)

// Kullanıcının Kayıt Olmasını Sağlayacak Olan Endpoint
// Min pass 7 cha
func SignUp(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-U-1"})
	}
	if len(user.Mail) == 0 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-U-2"})
	}
	if len(user.FullName) == 0 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-U-3"})
	}
	if len(user.Password) < 7 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-U-4"})
	}
	user.Password = helpers.HashPass(user.Password)
	err = database.DB.Db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-U-5"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})

}

// Kullanıcının Giriş Yapmasını Sağlayacak Olan Endpoint
func SignIn(c *fiber.Ctx) error {
	var user models.User
	var successUser models.SignIn
	var session models.Session
	err := c.BodyParser(&successUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-I-1"})
	}
	successUser.Password = helpers.HashPass(successUser.Password)
	err = database.DB.Db.Where("password=? and mail=?", successUser.Password, successUser.Mail).First(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-I-2"})
	}
	session.UserID = user.ID
	token, err := middleware.GenerateToken(user.Mail)
	session.Token = token
	err = database.DB.Db.Create(&session).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : S-I-3"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})

}

// Kullanıcın Şifresini Güncellemesini Sağlayacak Olan Enpoint
// Min Pass 7 Cha
func UpdatePassword(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-1"})
	}
	var updatepass models.UpdatePassword
	err := c.BodyParser(&updatepass)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-2"})
	}
	updatepass.OldPass = helpers.HashPass(updatepass.OldPass)
	if updatepass.OldPass != user.Password {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-3"})
	}
	if len(updatepass.NewPass1) < 7 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-4"})
	}
	if updatepass.NewPass1 != updatepass.NewPass2 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-5"})
	}
	user.Password = updatepass.NewPass1
	user.Password = helpers.HashPass(user.Password)
	err = database.DB.Db.Updates(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-6"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})
}

func UpdateAccount(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-A-1"})
	}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-A-2"})
	}
	err = database.DB.Db.Updates(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-A-3"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})
}

func DeleteAccount(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "D-A-1"})
	}
	id := c.Params("id")
	err := database.DB.Db.Where("id=?", id).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "D-A-2"})
	}
	err = database.DB.Db.Delete(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "D-A-3"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})

}

func LogOut(c *fiber.Ctx) error {
	sessionUser, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "user not found in context"})
	}
	err := database.DB.Db.Exec("DELETE FROM sessions WHERE user_id = ?", sessionUser.ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error delete session step"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "logout successfully"})
}

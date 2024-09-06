package controllers

import (
	"about/config"
	"about/database"
	"about/models"

	"github.com/gofiber/fiber/v2"
)

// Hakkımızda Oluştur
func CreateAboutSector(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : C-A-S-1"})
	}

	var about models.About
	err := c.BodyParser(&about)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : C-A-S-2"})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : C-A-S-3"})
	}

	fileBytes, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : C-A-S-4"})
	}
	defer fileBytes.Close()

	imageBytes := make([]byte, file.Size)
	_, err = fileBytes.Read(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : C-A-S-5"})
	}

	id, url, err := config.CloudConnect(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : C-A-S-6"})
	}

	about.Image = &id
	about.ImageURL = &url

	title := c.FormValue("title")
	text := c.FormValue("text")
	if len(title) != 0 {
		about.Title = title
	}

	if len(text) != 0 {
		about.Text = text
	}

	err = database.DB.Db.Create(&about).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : C-A-S-7"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Success", "data": about})

}

// Hakkımızda Kısmını Güncelle id url olarak alınacak
func UpdateAboutSector(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-1"})
	}
	id := c.Params("id")
	var about models.About
	err := database.DB.Db.Where("id=?", id).Find(&about).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-2"})
	}
	err = c.BodyParser(&about)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-3"})
	}
	file, err := c.FormFile("image")
	if err == nil {
		fileBytes, err := file.Open()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-4"})
		}
		defer fileBytes.Close()
		imageBytes := make([]byte, file.Size)
		_, err = fileBytes.Read(imageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-5"})
		}
		id, url, err := config.CloudConnect(imageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-6"})
		}
		about.Image = &id
		about.ImageURL = &url
	}
	title := c.FormValue("title")
	if len(title) != 0 {
		about.Title = title
	}
	text := c.FormValue("text")
	if len(text) != 0 {
		about.Text = text
	}
	err = database.DB.Db.Save(&about).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-7"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Successfully updated", "data": about})
}

// hakkımızda kısmında bir sutunu sil id url olarak alınacak
func DeleteAboutSector(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "ERROR: D-A-S-1"})
	}
	var about models.About
	id := c.Params("id")
	err := database.DB.Db.Where("id=?", id).First(&about).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-2"})
	}
	about.IsActive = false
	err = database.DB.Db.Updates(&about).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR: U-A-S-3"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Success"})

}

// Bütün hakkımızda kısmını getir sadece url olarak
func GetAllSector(c *fiber.Ctx) error {
	var about []models.About
	err := database.DB.Db.Where("is_active = ?", true).Order("id DESC").Find(&about).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "ERROR: G-A-S-1"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Success", "data": about})
}

// id ye göre hakkımızda kısmını getir
func GetIDBySector(c *fiber.Ctx) error {
	var about models.About
	id := c.Params("id")
	err := database.DB.Db.Where("id=?", id).First(&about).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : G-I-B-S"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": about})

}














//===================================İSTEĞE BAĞLI===================================
//Dinamik olarak URL oluşturup başka sayfaya atamasını sağlayacak olan endpoint

func CreateURL(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : C-U-1"})
	}
	var differentpageabout models.DiffrentPageAbout
	err := c.BodyParser(&differentpageabout)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : C-U-2"})
	}
	if len(differentpageabout.PageName) == 0 {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : C-U-3"})
	}
	err = database.DB.Db.Create(&differentpageabout).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : C-U-4"})
	}
	return c.Status(500).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": differentpageabout})
}

func UpdateURL(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-U-1"})
	}
	var diffrentpageabout models.DiffrentPageAbout
	err := c.BodyParser(&diffrentpageabout)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-U-2"})
	}
	id := c.Params("id")
	err = database.DB.Db.Where("id=?", id).First(&diffrentpageabout).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-U-3"})
	}
	err = database.DB.Db.Updates(&diffrentpageabout).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-U-4"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": diffrentpageabout})

}

func DeleteURL(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-U-1"})
	}
	var differentpageabout models.DiffrentPageAbout
	err := c.BodyParser(&differentpageabout)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-U-2"})
	}
	id := c.Params("id")
	err = database.DB.Db.Where("id=?", id).First(&differentpageabout).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-U-3"})
	}
	differentpageabout.IsActive = false
	err = database.DB.Db.Updates(&differentpageabout).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-U-4"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": differentpageabout})

}

func CreatePageInformation(c *fiber.Ctx) error {
	return nil
}

func UpdatePageInformation(c *fiber.Ctx) error {
	return nil
}

func DeletePageInformation(c *fiber.Ctx) error {
	return nil
}

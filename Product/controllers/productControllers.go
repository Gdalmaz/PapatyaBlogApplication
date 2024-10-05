package controllers

import (
	"product/config"
	"product/database"
	"product/models"

	"github.com/gofiber/fiber/v2"
)

func AddProduct(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : A-P-1"})
	}
	var product models.Product
	productname := c.FormValue("productname")
	productdescription := c.FormValue("productdescription")

	if len(productname) != 0 {
		product.ProductName = productname
	}
	if len(productdescription) != 0 {
		product.ProductDescription = productdescription
	}

	product.UserID = user.ID

	// Resmi yükleme işlemi
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-2"})
	}

	fileBytes, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-3"})
	}
	defer fileBytes.Close()

	imageBytes := make([]byte, file.Size)
	_, err = fileBytes.Read(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-4"})
	}

	id, url, err := config.CloudConnect(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-5"})
	}
	product.Image = id
	product.ImageUrl = url

	// Ürünü kaydetme
	err = database.DB.Db.Create(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : A-P-6"})
	}

	// Ekstra resim yükleme
	extraFile, err := c.FormFile("extra_image")
	if err == nil {
		extraFileBytes, err := extraFile.Open()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-EXTRA-1"})
		}
		defer extraFileBytes.Close()

		extraImageBytes := make([]byte, extraFile.Size)
		_, err = extraFileBytes.Read(extraImageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-EXTRA-2"})
		}

		extraId, extraUrl, err := config.CloudConnect(extraImageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-EXTRA-3"})
		}

		extraImage := models.ExtraImage{
			ProductID:       product.ID,
			ProductImage:    extraId,
			ProductImageUrl: extraUrl,
		}
		err = database.DB.Db.Create(&extraImage).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-P-EXTRA-4"})
		}
	}

	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Product added successfully"})
}

func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product models.Product
	err := database.DB.Db.First(&product, productID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "Product not found"})
	}
	user, ok := c.Locals("user").(models.User)
	if !ok || product.UserID != user.ID {
		return c.Status(403).JSON(fiber.Map{"Status": "Error", "Message": "Unauthorized"})
	}

	productname := c.FormValue("productname")
	productdescription := c.FormValue("productdescription")
	if len(productname) != 0 {
		product.ProductName = productname
	}
	if len(productdescription) != 0 {
		product.ProductDescription = productdescription
	}

	file, err := c.FormFile("image")
	if err == nil {
		fileBytes, err := file.Open()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-P-2"})
		}
		defer fileBytes.Close()

		imageBytes := make([]byte, file.Size)
		_, err = fileBytes.Read(imageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-P-3"})
		}

		id, url, err := config.CloudConnect(imageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-P-5"})
		}
		product.Image = id
		product.ImageUrl = url
	}

	extraFile, err := c.FormFile("extra_image")
	if err == nil {
		extraFileBytes, err := extraFile.Open()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-P-EXTRA-1"})
		}
		defer extraFileBytes.Close()

		extraImageBytes := make([]byte, extraFile.Size)
		_, err = extraFileBytes.Read(extraImageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-P-EXTRA-2"})
		}

		extraId, extraUrl, err := config.CloudConnect(extraImageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-P-EXTRA-3"})
		}

		var extraImage models.ExtraImage
		err = database.DB.Db.Where("product_id = ?", product.ID).First(&extraImage).Error
		if err != nil {
			extraImage = models.ExtraImage{
				ProductID:       product.ID,
				ProductImage:    extraId,
				ProductImageUrl: extraUrl,
			}
			err = database.DB.Db.Create(&extraImage).Error
		} else {
			extraImage.ProductImage = extraId
			extraImage.ProductImageUrl = extraUrl
			err = database.DB.Db.Save(&extraImage).Error
		}

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-P-EXTRA-4"})
		}
	}

	err = database.DB.Db.Save(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-6"})
	}

	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Product updated successfully"})
}

func DeleteProduct(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR: D-P-1"})
	}
	var product models.Product
	id := c.Params("id")
	err := database.DB.Db.Where("id=?", id).First(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-P-2"})
	}
	if user.ID != product.UserID {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-P-3"})
	}
	product.IsActive = false
	err = database.DB.Db.Save(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "Error : D-P-4"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})
}

// bu endpoint silinen product ları geri alacak
func RebinProduct(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : R-P-1"})
	}
	var product models.Product
	id := c.Params("id")
	err := database.DB.Db.Where("id=?", id).First(&product).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : R-P-2"})
	}
	product.IsActive = true
	err = database.DB.Db.Save(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : R-P-3"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})
}

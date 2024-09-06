package controllers

import (
	"product/config"
	"product/database"
	"product/models"

	"github.com/gofiber/fiber/v2"
)

func AddProduct(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : A-P-1"})
	}
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-R-2"})

	}
	productname := c.FormValue("productname")
	if len(productname) != 0 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : A-P-3"})
	}
	productdescription := c.FormValue("productdescription")
	if len(productdescription) != 0 {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : A-P-4"})
	}
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-R-5"})
	}

	fileBytes, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-R-6"})
	}

	defer fileBytes.Close()

	imageBytes := make([]byte, file.Size)
	_, err = fileBytes.Read(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-R-7"})
	}

	id, url, err := config.CloudConnect(imageBytes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-R-8"})
	}
	product.ProductImage = id
	product.ProductImageURL = url
	err = database.DB.Db.Create(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : A-R-9"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Success", "data": product})

}

func UpdateProduct(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-1"})
	}

	productID := c.Params("id")
	if productID == "" {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-2"})
	}

	var product models.Product

	if err := database.DB.Db.First(&product, productID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-3"})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : U-P-4"})
	}

	productname := c.FormValue("productname")
	if len(productname) != 0 {
		product.ProductName = productname
	}

	productdescription := c.FormValue("productdescription")
	if len(productdescription) != 0 {
		product.ProductDescription = productdescription
	}

	file, err := c.FormFile("image")
	if err == nil {
		fileBytes, err := file.Open()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-R-5"})
		}
		defer fileBytes.Close()

		imageBytes := make([]byte, file.Size)
		_, err = fileBytes.Read(imageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-R-6"})
		}

		id, url, err := config.CloudConnect(imageBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-R-7"})
		}
		product.ProductImage = id
		product.ProductImageURL = url
	}

	if err := database.DB.Db.Save(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "ERROR : U-R-8"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Product updated successfully", "data": product})
}

// Burada silmiyoruz sadece durumunu false yapıyoruz
func DeleteProduct(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(models.User)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-P-1"})
	}
	var product models.Product
	id := c.Params("id")
	err := database.DB.Db.Where("id=?", id).First(&product).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-P-2"})
	}
	product.IsActive = false
	err = database.DB.Db.Updates(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "ERROR : D-P-3"})
	}
	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success"})
}


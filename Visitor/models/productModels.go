package models

type Product struct {
	ID                 int    `json:"id"`
	ProductName        string `json:"productname"`
	ProductDescription string `json:"productdescription"`
	ProductImage       string `json:"productimage"`
	ProductImageURL    string `json:"productimageurl"`
	IsActive           bool   `gorm:"default:true"`
}

type ExtraImage struct {
	Product       Product `json:"ProductID"`
	ProductID     int     `json:"productid"`
	ExtraImage    string  `json:"extraimage"`
	ExtraImageURL string  `json:"extraimageurl"`
	IsActive      bool    `gorm:"default:true"`
}

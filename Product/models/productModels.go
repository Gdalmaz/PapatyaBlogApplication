package models

type Product struct {
	ID                 int          `json:"id"`
	User               User         `gorm:"foreignKey:UserID"` // ForeignKey tanımı
	UserID             int          `json:"user_id"`
	ProductName        string       `json:"product_name"`
	ProductDescription string       `json:"product_description"`
	ImageUrl           string       `json:"image_url"`
	Image              string       `json:"image"`
	IsActive           bool         `gorm:"default:true" json:"is_active"`
	ExtraImages        []ExtraImage `gorm:"foreignKey:ProductID" json:"extra_images"` // İlişki tanımı
}

type ExtraImage struct {
	Product         Product `gorm:"foreignKey:ProductID"` // ForeignKey tanımı düzeltildi
	ProductID       int     `json:"product_id"`
	ProductImageUrl string  `json:"product_image_url"`
	ProductImage    string  `json:"product_image"`
	IsActive        bool    `gorm:"default:true" json:"is_active"`
}

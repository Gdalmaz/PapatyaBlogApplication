package models

type About struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Text     string  `json:"text"`
	Image    *string `json:"image"`
	ImageURL *string `json:"imageurl"`
	IsActive bool    `gorm:"default:true"`
}

//BU KISIMLAR EXTRA EĞER MÜŞTERİ FAZLADAN URL AÇMAK İSTERSE DİYE TASARLANDI

type DiffrentPageAbout struct {
	ID       int    `json:"id"`
	PageName string `json:"pagename"`
	IsActive bool   `gorm:"default:true"`
}

type PageInformation struct {
	DiffrentPageAbout   DiffrentPageAbout `gorm:"foreignKey:DiffrentPageAboutID"`
	DiffrentPageAboutID int               `json:"diffrentpageaboutid"`
	Title               string            `json:"title"`
	Content             string            `json:"content"`
	Image               *string           `json:"image"`
	ImageURL            *string           `json:"imageurl"`
	IsActive            bool              `gorm:"default:true"`
}

package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type Session struct {
	UserID int    `gorm:"primaryKey;autoIncrement`
	Token  string `json:"token"`
}

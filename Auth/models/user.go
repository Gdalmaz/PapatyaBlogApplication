package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type UpdatePassword struct {
	OldPass  string `json:"oldpass"`
	NewPass1 string `json:"newpass1"`
	NewPass2 string `json:"newpass2"`
}

type SignIn struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

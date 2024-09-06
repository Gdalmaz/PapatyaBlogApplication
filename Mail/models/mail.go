package models

type Mail struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	WhoSend string `json:"whosend"`
}

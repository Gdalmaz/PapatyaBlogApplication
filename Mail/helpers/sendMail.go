package helpers

import (
	"log"
	"mail/config"
)

func SendMail(mail string, text string) error {
	mainMail := "gokhandalmzz@gmail.com"
	err := config.RabbitMqPublish([]byte(text), mail)
	if err != nil {
		log.Println("Publish hatası")
		return err
	}
	err = config.RabbitMqConsume(mail, mainMail)
	if err != nil {
		log.Println("Consume hatası")
		return err
	}
	return nil
}

package helpers

import "mail/config"

func SendMail(mail string, text string) error {
	mainMail := "gokhandalmzz@gmail.com"
	err := config.RabbitMqPublish([]byte(text), mail)
	if err != nil {
		return err
	}
	err = config.RabbitMqConsume(mail, mainMail)
	if err != nil {
		return err
	}
	return nil
}

package mail

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"

	config "tritan.dev/config"
)

func SendMail(address string, ctx *fiber.Ctx) {
	config := ctx.Locals("config").(*config.AppConfig)

	m := mail.NewMessage()
	m.SetHeader("From", config.Mail.FromAddress)
	m.SetHeader("To", address)
	m.SetHeader("Subject", "Please verify your email!")
	m.SetBody("text/html", "Hi there! <br> <br> We're glad you signed up for our services, please click the link below to verify your account.")

	d := mail.NewDialer(config.Mail.MailServer, config.Mail.Port, config.Mail.Username, config.Mail.Password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

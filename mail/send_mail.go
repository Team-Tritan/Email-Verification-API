package mail

import (
	"fmt"
	"strings"
  	"text/template"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"
	

	"tritan.dev/config"
)

func SendMail(email string, userkey string, ctx *fiber.Ctx) {
	config := ctx.Locals("config").(*config.AppConfig)

	str := config.Baseurl + "/verify/" + userkey

	m := mail.NewMessage()
	m.SetHeader("From", config.Mail.FromAddress)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Please verify your email!")
	m.SetBody("text/html", formatTemplate(`<body spellcheck="false"><div class="editor"><div class="editable-area"><style>@import url(https://fonts.googleapis.com/css2?family=Fira+Sans:ital,wght@0,400;0,700;1,400;1,700&display=swap);a:active,a:hover,a:link,a:visited{color:#706fd3}body{font-family:"Work Sans",sans-serif;font-size:16px;line-height:1.4;padding:0;margin:0 2rem;box-sizing:border-box}.editor{max-width:1200px;margin:2rem auto 2rem;position:relative}.tinymce:focus{border-radius:.5px;box-shadow:0 0 0 4px #fff,0 0 0 7px #706fd333;outline:0}</style><table border="0"style="background-color:#f2f2f2;width:100%"><tr><td align="center"cellpadding="16px"style="padding:16px"><table border="0"style="max-width:600px;width:100%;background-color:#fff;border-top-left-radius:16px;border-top-right-radius:16px;padding-top:32px;padding-left:32px;padding-right:32px"width="100%"><tr><td width="64px"><img alt="Logo"height="64"src="https://www.tiny.cloud/storage/codepens/email-marketing/vcf-logo-2.png"width="64"><td align="right"><p style="font-family:'Fira Sans',sans-serif;font-size:17px;font-weight:700"><strong>Email Verification</strong></table><table border="0"style="max-width:600px;width:100%;background-color:#fff;padding-left:32px;padding-right:32px;padding-bottom:32px"width="100%"><tr><td><div class="mce-content-body tinymce"style='font-family:"Fira Sans",sans-serif;color:#222;font-size:15px;line-height:1.5;position:relative'contenteditable="true"id="tinymce"spellcheck="false"><h1 data-mce-style="font-size: 32px; font-weight: bold;"style="font-size:32px;font-weight:700">Thank you for registering for our service!</h1><p>Please select the link below to verify that you own this email. If you did not request this, you may ignore this message.<p><a href="{{.}}"data-mce-href="{{.}}"data-mce-style="background-color: #706fd3; padding: 12px 16px; color: #ffffff; border-top-left-radius: 4px; border-top-right-radius: 4px; border-bottom-right-radius: 4px; border-bottom-left-radius: 4px; text-decoration: none; display: inline-block;"style="background-color:#706fd3;padding:12px 16px;color:#fff;border-top-left-radius:4px;border-top-right-radius:4px;border-bottom-right-radius:4px;border-bottom-left-radius:4px;text-decoration:none;display:inline-block">Verify me, baby!</a></div></table></div></div><div class="tox-anchorbar"></div><div class="tox-throbber"style="display:none"aria-hidden="true"></div><div class="tox tox-silver-sink tox-tinymce-aux"style="position:relative;display:none"></div><p style="color:grey">Made with 💜 by <a href="https://tritan.dev">Tritan Development</a>`, str))

	d := mail.NewDialer(config.Mail.MailServer, config.Mail.Port, config.Mail.Username, config.Mail.Password)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("%v", err)
		ctx.JSON(fiber.Map{
			"error":   true,
			"code":    500,
			"message": err,
			"email":   email,
		})
	}
}


func formatTemplate(s string, v interface{}) string {
   t, b := new(template.Template), new(strings.Builder)
   template.Must(t.Parse(s)).Execute(b, v)
   return b.String()
}
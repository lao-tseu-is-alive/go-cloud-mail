package main

import (
	"gopkg.in/gomail.v2"
)

const (
	VERSION             = "0.1.0"
	APP                 = "go-cloud-mail"
	defaultSMTPServer   = "smtp.gmail.com"
	defaultSMTPPort     = 587
	defaultSMTPUser     = "goeland.lausanne@gmail.com"
	defaultSMTPPassword = "defineYourOwn"
	defaultSubject      = "🄷🄴🄻🄻🄾 🄵🅁🄾🄼 🄶🄾🄻🄰🄽🄶 😎"
	defaultRecipient    = "lao.tseu.is.alive@gmail.com"
	defaultCC           = "carlos.gil@lausanne.ch"
)

func main() {

	msg := gomail.NewMessage()
	msg.SetHeader("From", defaultSMTPUser)
	msg.SetHeader("To", defaultRecipient, "cnaegele@lausanne.ch")
	msg.SetAddressHeader("Cc", defaultCC, "CG")
	msg.SetHeader("Subject", defaultSubject)
	msg.SetBody("text/html", "<h2>Bonjour !</h2> <b>Ce message a été envoyé par Goéland</b>")
	msg.Attach("./cat.jpg")

	n := gomail.NewDialer(defaultSMTPServer, defaultSMTPPort, defaultSMTPUser, defaultSMTPPassword)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}

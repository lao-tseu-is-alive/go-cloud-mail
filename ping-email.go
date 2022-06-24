package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/go-cloud-mail/pkg/config"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

const (
	VERSION           = "0.1.0"
	APP               = "go-cloud-mail"
	defaultSMTPServer = "smtp.gmail.com"
	defaultSMTPPort   = 587
	defaultSMTPUser   = "goeland.lausanne@gmail.com"
	defaultSubject    = "🄷🄴🄻🄻🄾 🄵🅁🄾🄼 🄶🄾🄻🄰🄽🄶 😎"
	defaultRecipient  = "lao.tseu.is.alive@gmail.com"
	defaultCC         = "carlos.gil@lausanne.ch"
)

func main() {

	l := log.New(os.Stdout, fmt.Sprintf("[%s]", APP), log.Ldate|log.Ltime|log.Lshortfile)
	l.Printf("INFO: 'Starting %s version:%s'\n", APP, VERSION)
	msg := gomail.NewMessage()
	msg.SetHeader("From", defaultSMTPUser)
	msg.SetHeader("To", defaultRecipient, "cnaegele@lausanne.ch")
	msg.SetAddressHeader("Cc", defaultCC, "CG")
	msg.SetHeader("Subject", defaultSubject)
	msg.SetBody("text/html", "<h2>Bonjour !</h2> <b>Ce message a été envoyé par Goéland</b>")
	msg.Attach("./cat.jpg")

	smtpConn, err := config.GetSmtpConnectInfoFromEnv(defaultSMTPServer, defaultSMTPPort, defaultSMTPUser, "")
	if err != nil {
		l.Fatalf("💥💥 ERROR: 'calling GetSmtpConnectInfoFromEnv got error: %v'\n", err)
	}
	l.Printf("INFO: 'mail.NewDialer(server: %s, port:%d, user: %s'\n", smtpConn.Server, smtpConn.Port, smtpConn.User)
	n := gomail.NewDialer(smtpConn.Server, smtpConn.Port, smtpConn.User, smtpConn.Password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		l.Fatalf("💥💥 ERROR: 'calling DialAndSend got error: %v'\n", err)
	}
	l.Printf("INFO: 'Ending %s version:%s'\n", APP, VERSION)
}

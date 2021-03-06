package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/go-cloud-mail/pkg/config"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	VERSION           = "0.1.0"
	APP               = "ping-email"
	defaultSMTPServer = "smtp.gmail.com"
	defaultSMTPPort   = 587
	defaultSMTPUser   = "your.user@gmail.com"
	//defaultSubject    = "π±πΈπ΄π½ππ΄π½ππ΄ πππ πΆπΎπ΄π»π°π½π³  π¦"
	defaultSubject   = "ππππ‘π©ππ‘π¨π π¦π¨π₯ ππ’ππππ‘π π¦"
	defaultRecipient = "lao.tseu.is.alive@gmail.com"
	defaultCC        = "carlos.gil@lausanne.ch"
	defaultTemplate  = "templates/hello.html"
)

func main() {

	l := log.New(os.Stdout, fmt.Sprintf("[%s]", APP), log.Ldate|log.Ltime|log.Lshortfile)
	l.Printf("INFO: 'Starting %s version:%s'\n", APP, VERSION)
	l.Printf("INFO: 'about to read email template : %s'\n", defaultTemplate)
	body, err := ioutil.ReadFile(defaultTemplate)
	if err != nil {
		l.Fatalf("π₯π₯ ERROR: 'problem doing ioutil.ReadFile(%s) got error: %v'\n", defaultTemplate, err)
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", defaultSMTPUser)
	msg.SetHeader("To", defaultRecipient, "cnaegele@lausanne.ch", "cnaegele62@gmail.com")
	msg.SetAddressHeader("Cc", defaultCC, "CG")
	msg.SetHeader("Subject", defaultSubject)
	msg.Embed("./templates/logo.png")
	msg.SetBody("text/html", string(body))
	msg.Attach("./cat.jpg")

	smtpConn, err := config.GetSmtpConnectInfoFromEnv(defaultSMTPServer, defaultSMTPPort, defaultSMTPUser, "")
	if err != nil {
		l.Fatalf("π₯π₯ ERROR: 'calling GetSmtpConnectInfoFromEnv got error: %v'\n", err)
	}
	if smtpConn.Password == "" {
		l.Fatal("π₯π₯ ERROR: 'SMTP_PASSWORD cannot be an empty string'")
	}
	l.Printf("INFO: 'mail.NewDialer(server: %s, port:%d, user: %s'\n", smtpConn.Server, smtpConn.Port, smtpConn.User)
	n := gomail.NewDialer(smtpConn.Server, smtpConn.Port, smtpConn.User, smtpConn.Password)

	l.Printf("INFO: 'calling DialAndSend(msg): will send message to %s'\n", defaultRecipient)
	if err := n.DialAndSend(msg); err != nil {
		l.Fatalf("π₯π₯ ERROR: 'calling DialAndSend got error: %v'\n", err)
	}
	l.Printf("INFO: 'Ending %s version:%s'\n", APP, VERSION)
}

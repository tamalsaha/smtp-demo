package main

import (
	"fmt"
	"net/smtp"
	"os"

	"gomodules.xyz/email"
)

func main() {
	e := email.NewEmail()
	e.From = "Go Code <gocode@appscode.com>"
	e.To = []string{"tamal@appscode.com"}
	e.Bcc = []string{"appscode@appscode.com"}
	e.Cc = []string{"saif@appscode.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")

	host := os.Getenv("MG_DOMAIN")
	username := os.Getenv("MG_SMTP_USERNAME")
	password := os.Getenv("MG_SMTP_PASSWORD")

	fmt.Println(host, username)

	err := e.Send("smtp.mailgun.org:587", smtp.PlainAuth("", username, password, "smtp.mailgun.org"))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

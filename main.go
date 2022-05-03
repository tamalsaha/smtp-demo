package main

import (
	"fmt"
	"net"
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
	e.AttachFile("/Users/tamal/go/src/github.com/tamalsaha/smtp-demo/main.go")
	e.AttachFile("/Users/tamal/go/src/github.com/tamalsaha/smtp-demo/go.mod")
	e.AttachFile("/Users/tamal/Downloads/AppsCode - Backend Engineer.pdf")

	addr := os.Getenv("SMTP_ADDRESS")
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	err = e.Send(addr, smtp.PlainAuth("", username, password, host))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

package main

import (
	"fmt"
    "bytes"
    "html/template"
	"github.com/benmanns/goworker"
    "gopkg.in/gomail.v1"
)

var doc bytes.Buffer

type SmtpSettings struct {
    Email    string
    Password    string
    Server string
    Port        int
}

func init() {
	goworker.Register("Mailer", MailWorker)
}

func MailWorker(queue string, args ...interface{}) error {
    smtp := &SmtpSettings{ "admin@gmail.com", "password", "smtp.gmail.com", 465 }

    tmpl, err := template.New("emailTemplate").Parse(`Hello, {{.}}`)
    if err != nil { panic(err) }
    err = tmpl.Execute(&doc, args[1])
    if err != nil { panic(err) }

    msg := gomail.NewMessage()
    msg.SetHeader("From", smtp.Email)
    msg.SetHeader("To", args[0].(string))
    msg.SetHeader("Subject", "Hello!")
    msg.SetBody("text/html", doc.String())

    mailer := gomail.NewMailer(smtp.Server, smtp.Email, smtp.Password, smtp.Port)
    if err := mailer.Send(msg); err != nil {
        panic(err)
    }
	fmt.Println("Email was sent!")
	return nil
}

func main(){
    if err := goworker.Work(); err != nil {
        fmt.Println("Error:", err)
    }
}
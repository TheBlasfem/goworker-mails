package main

import (
	"fmt"
    "bytes"
    "html/template"
	"github.com/benmanns/goworker"
    "gopkg.in/gomail.v1"
)

var doc bytes.Buffer

func init() {
	goworker.Register("Mailer", MailWorker)
}

func MailWorker(queue string, args ...interface{}) error {
    tmpl, err := template.New("emailTemplate").Parse(`Hello, {{.}}`)
    if err != nil { panic(err) }
    err = tmpl.Execute(&doc, args[0])
    if err != nil { panic(err) }

    msg := gomail.NewMessage()
    msg.SetHeader("From", args[0].(string))
    msg.SetHeader("To", args[1].(string))
    msg.SetHeader("Subject", "Hello!")
    msg.SetBody("text/html", doc.String())

    mailer := gomail.NewMailer("smtp.gmail.com", args[0].(string), "password", 465)
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
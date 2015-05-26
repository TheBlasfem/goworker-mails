package main

import (
	"fmt"
	"github.com/benmanns/goworker"
  "gopkg.in/gomail.v1"
)

func init() {
	goworker.Register("Hello", helloWorker)
}

func helloWorker(queue string, args ...interface{}) error {
  msg := gomail.NewMessage()
  msg.SetHeader("From", "sender@gmail.com")
  msg.SetHeader("To", "receiver@outlook.com")
  msg.SetHeader("Subject", "Hello!")
  msg.SetBody("text/html", "Hello <b>Julio</b>!")

  mailer := gomail.NewMailer("smtp.gmail.com", "sender@gmail.com", "password", 465)
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
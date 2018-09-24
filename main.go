package main

import (
	"eFuture/mail"
	"eFuture/rabbitmq"
)

func main() {
	rabbitmq.Push("Hello!")
	go func() {
		rabbitmq.Receive()
	}()
	println("try to send mail")
	mail.SendMail("hello xlui!", []string{"i@xlui.me"}, "I'm sending mail to you through golang.")
}
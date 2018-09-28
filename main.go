package main

import (
	"eFuture/rabbitmq"
	"time"
)

func main() {
	receive := time.Now().Add(time.Minute)
	rabbitmq.Push("hello, world!")
	rabbitmq.PushAtDate("Hello!", receive)
	rabbitmq.Receive()
}

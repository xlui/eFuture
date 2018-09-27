package main

import (
	"eFuture/rabbitmq"
	"time"
)

func main() {
	receive := time.Now().Add(time.Second * 5)
	rabbitmq.Push("hello, world!")
	rabbitmq.PushAtDate("Hello!", receive)
	rabbitmq.Receive()
}

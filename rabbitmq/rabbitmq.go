//package rabbitmq
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

const (
	exchange = "eFuture.msg.exchange"
	queue    = "eFuture.msg.queue"
	key		 = "queue_ex"
	mqUrl    = "amqp://user:user@localhost:5672/eFuture"
)

var (
	connection *amqp.Connection
	channel    *amqp.Channel
)

func main() {
	go func() {
		for {
			Push("hello, world!")
			time.Sleep(1 * time.Second)
		}
	}()
	Receive()
	fmt.Println("End!")
	close()
}

func connect() {
	connection, e := amqp.Dial(mqUrl)
	if e != nil {
		log.Fatalf("%s:%s", "Failed to connect to rabbitmq", e)
	}
	channel, e = connection.Channel()
	if e != nil {
		log.Fatalf("%s:%s", "Failed to open a channel", e)
	}
	channel.QueueBind(queue, key, exchange, false, nil)
}

func close() {
	channel.Close()
	connection.Close()
}

func Push(message string) {
	if channel == nil {
		connect()
	}
	channel.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}

func Receive() {
	if channel == nil {
		connect()
	}
	messages, e := channel.Consume(queue, "", true, false, false, false, nil)
	if e != nil {
		log.Fatalf("%s:%s", "", e)
	}

	forever := make(chan bool)
	go func() {
		for message := range messages {
			s := string(message.Body)
			fmt.Printf("Receive: %s -- %s\n", s, time.Now().Format(time.RFC3339))
		}
	}()
	fmt.Println("[*] Waiting for messages. To exist press CTRL+C")
	<-forever
}
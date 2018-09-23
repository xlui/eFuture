package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

const (
	exchange = "eFuture.msg.exchange"
	queue    = "eFuture.msg.queue"
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
}

func close() {
	channel.Close()
	connection.Close()
}

func Push(message string) {
	if channel == nil {
		connect()
	}
	channel.Publish(exchange, queue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}

func Receive() {
	if channel == nil {
		connect()
	}
	count := 0
	messages, e := channel.Consume(queue, "", true, false, false, false, nil)
	if e != nil {
		log.Fatalf("%s:%s", "", e)
	}

	forever := make(chan bool)
	go func() {
		for message := range messages {
			s := string(message.Body)
			count++
			fmt.Printf("Receive: %s -- %d\n", s, count)
		}
	}()
	fmt.Println("[*] Waiting for messages. To exist press CTRL+C")
	<-forever
}

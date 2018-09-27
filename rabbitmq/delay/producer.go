package main

import (
	"eFuture/rabbitmq/common"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main() {
	connection, e := amqp.Dial("amqp://user:user@localhost:5672")
	common.FailOnError(e, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, e := connection.Channel()
	common.FailOnError(e, "Failed to open a channel!")
	defer channel.Close()

	body := bodyForm(os.Args)
	e = channel.Publish(
		"",
		"test_delay",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:  "5000", // 5秒过期
		},
	)
	common.FailOnError(e, "Failed to publish a message!")
	log.Printf("[x] Sent %s", body)
}

func bodyForm(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

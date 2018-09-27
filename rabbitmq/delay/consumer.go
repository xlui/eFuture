package main

import (
	"eFuture/rabbitmq/common"
	"github.com/streadway/amqp"
	"log"
)


func main() {
	connection, e := amqp.Dial("amqp://user:user@localhost:5672")
	common.FailOnError(e, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, e := connection.Channel()
	common.FailOnError(e, "Failed to open a channel!")
	defer channel.Close()

	e = channel.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	common.FailOnError(e, "Failed to declare an exchange")

	queue, e := channel.QueueDeclare(
		"test_logs",
		false,
		false,
		true,
		false,
		nil,
	)
	common.FailOnError(e, "Failed to declare a queue!")

	_, e = channel.QueueDeclare(
		"test_delay",
		false,
		false,
		true,
		false,
		amqp.Table{
			"x-dead-letter-exchange": "logs",
		},
	)
	common.FailOnError(e, "Failed to declare a delay queue")

	e = channel.QueueBind(queue.Name, "", "logs", false, nil)
	common.FailOnError(e, "Failed to bind a queue!")

	messages, e := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	common.FailOnError(e, "Failed to register a consumer!")

	forever := make(chan bool)
	go func() {
		for message := range messages {
			log.Printf("[x] %s", message.Body)
		}
	}()
	log.Printf("[*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

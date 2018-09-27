package rabbitmq

import (
	"eFuture/common"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

const (
	exchange   = "eFuture.msg.exchange"
	queue      = "eFuture.msg.queue"
	delayQueue = "eFuture.msg.future"
	mqUrl      = "amqp://user:user@localhost:5672/eFuture"
)

var connection *amqp.Connection
var channel *amqp.Channel

func init() {
	connect()
}

func connect() {
	var e error
	// connect
	connection, e = amqp.Dial(mqUrl)
	common.FailOnError(e, "Failed to connect to rabbitmq!")
	// open channel
	channel, e = connection.Channel()
	common.FailOnError(e, "Failed to open a channel!")
	channel.QueueBind(queue, delayQueue, exchange, false, nil)
}

//noinspection GoUnusedFunction,GoReservedWordUsedAsName
func close() {
	channel.Close()
	connection.Close()
}

func Push(message string) {
	channel.Publish(exchange, delayQueue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
}

func PushAtDate(message string, date time.Time) {
	// millisecond
	duration := date.Sub(time.Now()).Seconds() * 1000
	// to-string
	expire := strconv.Itoa(int(duration))
	log.Println(expire)
	e := channel.Publish(
		"",
		delayQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			Expiration:  expire, // 5秒过期
		},
	)
	common.FailOnError(e, "Failed to send message!")
}

func Receive() {
	messages, e := channel.Consume(queue, "", true, false, false, false, nil)
	common.FailOnError(e, "")
	forever := make(chan bool)
	go func() {
		for message := range messages {
			log.Println(string(message.Body))
		}
	}()
	fmt.Println("[*] Waiting for messages. To exist press CTRL+C")
	<-forever
}

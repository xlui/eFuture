package redis

import (
	"eFuture/common"
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	_, e := client.Ping().Result()
	common.FailOnError(e, "Failed to connect to redis")
}

func TestQueueDelete(t *testing.T) {
	result := client.Get(QUEUE_KEY).Val()
	if len(result) != 0 {
		log.Fatalln("Failed to delete exist queue!")
	}
}

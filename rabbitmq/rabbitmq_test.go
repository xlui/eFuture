package rabbitmq

import (
	"log"
	"testing"
)

func TestChannel(t *testing.T) {
	if channel == nil {
		log.Fatalln("Failed to open channel!")
		return
	}
	defer channel.Close()
}

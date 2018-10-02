package main

import (
	"eFuture/redis"
	"log"
	"time"
)

func main() {
	date1 := time.Now().Add(time.Second * 24)
	date2 := time.Now().Add(time.Second * 5)
	date3 := time.Now().Add(time.Second * 30)
	log.Println("Push data 1")
	redis.Push("hello 1", date1)
	log.Println("Push data 2")
	redis.Push("hello 2", date2)
	log.Println("Push data 3")
	redis.Push("hello 3", date3)
	for {
		time.Sleep(1 * time.Second)
		if b, d := redis.Pop(); b {
			log.Println(d)
		}
	}
}

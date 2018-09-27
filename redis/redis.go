package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, e := client.Ping().Result()
	fmt.Println(pong, e)

	e = client.Set("key", "val", 0).Err()
	if e != nil {
		panic(e)
	}

	val, e := client.Get("key").Result()
	if e != nil {
		panic(e)
	}
	fmt.Println("key:", val)

	val2, e := client.Get("key2").Result()
	if e == redis.Nil {
		fmt.Println("key2 does not exist!")
	} else if e != nil {
		panic(e)
	} else {
		fmt.Println("key2:", val2)
	}
}

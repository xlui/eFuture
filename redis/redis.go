// Redis delayed task implement
package redis

import (
	"eFuture/common"
	"eFuture/config"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

const QUEUE_KEY = "mail_queue"

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     config.Configuration.RedisAddress,
		Password: config.Configuration.RedisPassword,
		DB:       config.Configuration.RedisDB,
	})
	_, e := client.Ping().Result()
	common.FailOnError(e, "Failed to connect to redis")
	client.Del(QUEUE_KEY)
}

func Push(message string, date time.Time) {
	id := uuid.NewV4().String()
	client.Set(id, message, 0)
	client.ZAdd(QUEUE_KEY, redis.Z{
		Score:  float64(date.Unix()),
		Member: id,
	})
}

func Pop() (bool, string) {
	task := client.ZRange(QUEUE_KEY, 0, 0).Val()
	if len(task) == 0 {
		return false, ""
	}
	id := task[0]
	current := int64(client.ZScore(QUEUE_KEY, id).Val())
	if current <= time.Now().Unix() {
		data, _ := client.Get(id).Result()
		pipeline := client.Pipeline()
		pipeline.ZRem(QUEUE_KEY, id)
		pipeline.Del(id)
		_, e := pipeline.Exec()
		common.FailOnError(e, "Failed to exec delete")
		return true, data
	} else {
		return false, ""
	}
}

func main() {
	date1 := time.Now().Add(time.Second * 1)
	date2 := time.Now().Add(time.Second * 2)
	date3 := time.Now().Add(time.Second * 3)
	log.Println("Push data 1")
	Push("hello 1", date1)
	log.Println("Push data 2")
	Push("hello 2", date2)
	log.Println("Push data 3")
	Push("hello 3", date3)
	for {
		time.Sleep(1 * time.Second)
		if b, d := Pop(); b {
			log.Println(d)
		}
	}

}

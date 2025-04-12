package queue

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // настрой по .env позже
})

func PushEvent(event interface{}) {
	data, _ := json.Marshal(event)

	err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "event_stream",
		Values: map[string]interface{}{
			"data": string(data),
			"ts":   time.Now().Unix(),
		},
	}).Err()

	if err != nil {
		log.Printf("Failed to push event: %v\n", err)
	}
}

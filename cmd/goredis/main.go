package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

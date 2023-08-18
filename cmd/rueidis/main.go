package main

import (
	"context"
	"github.com/redis/rueidis"
	"log"
)

func main() {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	// ping
	if err := client.Do(ctx, client.B().Ping().Build()).Error(); err != nil {
		log.Fatal(err)
	}
}

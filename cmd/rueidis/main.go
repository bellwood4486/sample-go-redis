package main

import (
	"context"
	"github.com/redis/rueidis"
	"log"
)

func main() {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{
			InitAddress: []string{"127.0.0.1:6379"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := test(client); err != nil {
		log.Fatal(err)
	}
}

func test(client rueidis.Client) error {
	ctx := context.Background()
	var err error
	// PING
	err = client.Do(ctx, client.B().Ping().Build()).Error()
	if err != nil {
		return err
	}
	// SET
	err = client.Do(ctx, client.B().Set().Key("key").Value("value").Build()).Error()
	if err != nil {
		return err
	}
	// GET
	v, err := client.Do(ctx, client.B().Get().Key("key").Build()).ToString()
	if err != nil {
		return err
	}
	log.Printf("GET key: %s", v)

	return nil
}

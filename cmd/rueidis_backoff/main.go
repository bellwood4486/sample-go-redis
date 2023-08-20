package main

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/redis/rueidis"
	"log"
	"time"
)

func main() {
	client, err := rueidis.NewClient(
		rueidis.ClientOption{
			InitAddress:  []string{"127.0.0.1:6379"},
			DisableRetry: true,
		},
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed new client: %w", err))
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
		return fmt.Errorf("ping failed: %w", err)
	}
	// SET
	err = client.Do(ctx, client.B().Set().Key("key").Value("value").Build()).Error()
	if err != nil {
		return fmt.Errorf("set failed: %w", err)
	}
	// GET
	cmd := client.B().Get().Key("key").Build()
	v, err := do(ctx, client, cmd, withMaxElapsedTime(1*time.Minute)).ToString()
	if err != nil {
		return fmt.Errorf("get failed: %w", err)
	}
	log.Printf("GET key: %s", v)

	return nil
}

type retryOption func(b *backoff.ExponentialBackOff)

func withMaxElapsedTime(d time.Duration) retryOption {
	return func(b *backoff.ExponentialBackOff) {
		b.MaxElapsedTime = d
	}
}

func do(ctx context.Context, client rueidis.Client, cmd rueidis.Completed, opts ...retryOption) rueidis.RedisResult {
	// 自前のリトライでコマンドを使い回すのでPinしておく。
	cmd.Pin()
	var result rueidis.RedisResult
	op := func() error {
		log.Println("do cmd...")
		result = client.Do(ctx, cmd)
		return result.Error()
	}
	// setup backoff
	b := backoff.NewExponentialBackOff()
	for _, opt := range opts {
		opt(b)
	}

	_ = backoff.Retry(op, b)

	return result
}

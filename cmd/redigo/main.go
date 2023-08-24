package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	_, err = c.Do("SET", "hello", "world")
	if err != nil {
		log.Fatal(err)
	}
	s, err := redis.String(c.Do("GET", "hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", s)
}

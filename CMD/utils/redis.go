package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client  *redis.Client

func Redis() {
	fmt.Println("Testing Golang Redis")

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// pong, err := client.Ping(client.ping().Result())
	// fmt.Println(pong, err)

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	if err != nil {

		fmt.Println("ERROR", err)
	} else {

		fmt.Print("RESULT .........................", pong)
	}
}

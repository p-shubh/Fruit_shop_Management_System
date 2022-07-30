package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func Redis() {
	fmt.Println("Testing Golang Redis")

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// pong, err := client.Ping(client.ping().Result())
	// fmt.Println(pong, err)

	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)

	if err != nil {

		fmt.Println("ERROR", err)
	} else {

		fmt.Print("RESULT .........................", pong)
	}
}

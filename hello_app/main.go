package main

import (
	"fmt"
	"net/http"
	"time"

	redis "github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var cache *redis.Client

func getHitCount() int64 {
	retries := 5
	for {
		result, err := cache.Incr(context.Background(), "hits").Result()
		if err == nil {
			return result
		}

		if err == redis.Nil {
			_, err := cache.Set(context.Background(), "hits", 0, 0).Result()
			if err != nil {
				panic(err)
			}
		}

		if retries == 0 {
			panic(err)
		}

		retries--
		time.Sleep(500 * time.Millisecond)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	count := getHitCount()
	fmt.Fprintf(w, "Hello World! I have been seen %d times.\n", count)
}

func main() {
	cache = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // No password
		DB:       0,  // Default DB
	})

	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on :8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		panic(err)
	}
}

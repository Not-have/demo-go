package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	} else {
		fmt.Println("Connected to Redis successfully!")
	}

	// Set a key, 0 代表没有过期时间
	err01 := rdb.Set(ctx, "name", "呵呵呵", 0).Err()

	if err01 != nil {
		fmt.Println("Failed to set key:", err01)
		return
	}

	result := rdb.Get(ctx, "name").Val()
	fmt.Println(result)
}

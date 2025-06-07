package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var cache sync.Map

func main() {
	wg := sync.WaitGroup{}

	userIDs := []string{"user1", "user2", "user1", "user3", "user2", "user1"}

	for _, userID := range userIDs {
		wg.Add(1)

		go func(userID string) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.IntN(5)) * time.Second)

			res := GetOrCompute(userID, func() string {
				return computation(userID)
			})

			fmt.Println(res)
		}(userID)
	}

	wg.Wait()
}

func GetOrCompute(key string, computeFunc func() string) string {
	if val, ok := cache.Load(key); ok {
		fmt.Printf("from cache by %s\n", key)

		return val.(string)
	}

	newVal := computeFunc()

	cache.Store(key, newVal)

	return newVal
}

func computation(userID string) string {
	time.Sleep(1 * time.Second)

	return fmt.Sprintf("Computation for user %s", userID)
}

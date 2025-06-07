package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func LongCalculation(n int) int {
	secondsToSleep := rand.Float64() * float64(n)
	time.Sleep(time.Duration(secondsToSleep))
	return n + 1
}

type Cache struct {
	mu   sync.RWMutex
	data map[int]int
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]int),
	}
}

var cache = NewCache()

func CachedLongCalculation(n int) int {
	//var mu sync.Mutex // используем RWMutex?

	cache.mu.RLock() // исправили на RLock()
	found, ok := cache.data[n]
	cache.mu.RUnlock() // исправили на RUnlock()

	if !ok {
		value := LongCalculation(n)
		cache.mu.Lock()
		cache.data[n] = value
		cache.mu.Unlock()
		return value
	}

	//mu.Unlock() // лишний unlock?

	return found
}

func main() {
	nums := []int{5, 10, 22}

	for _, num := range nums {
		// async?
		val := CachedLongCalculation(num)
		fmt.Printf("LongCalculation(%d) = %d\n", num, val)
	}
}

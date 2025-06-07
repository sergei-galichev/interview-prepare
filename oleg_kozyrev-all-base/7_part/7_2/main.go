// Напишите реализацию In-Memory кэша

package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type Cache interface {
	Set(k, v string)
	Get(k string) (string, bool)
}

type Shard struct {
	mu   sync.RWMutex
	data map[string]string
}

type MyCache struct {
	shards []*Shard
}

func NewMyCache(shardCount int) *MyCache {
	shards := make([]*Shard, shardCount)

	for i := 0; i < shardCount; i++ {
		shards[i] = &Shard{
			data: make(map[string]string),
		}
	}

	return &MyCache{
		shards: shards,
	}
}

func (c *MyCache) Set(k, v string) {
	shard := c.getShard(k)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[k] = v
}

func (c *MyCache) Get(k string) (string, bool) {
	shard := c.getShard(k)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	result, ok := shard.data[k]

	return result, ok
}

func (c *MyCache) getShard(k string) *Shard {
	hasher := fnv.New32a()
	_, err := hasher.Write([]byte(k))
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	hash := hasher.Sum32()

	return c.shards[hash%uint32(len(c.shards))]
}

func main() {
	cache := NewMyCache(100)
	cache.Set("salary", "500000")

	value, ok := cache.Get("salary")
	if !ok {
		fmt.Println("no shards by key 'salary'")
		return
	}

	fmt.Println("Value by key 'salary':", value)

}

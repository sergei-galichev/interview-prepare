package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data map[string]string
	m    sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]string),
	}
}

func (c *ConcurrentMap) GetOrCreate(key string, value string) string {
	c.m.RLock()
	val, ok := c.data[key]
	c.m.RUnlock()

	if ok {
		return val
	}

	c.m.Lock()
	defer c.m.Unlock()
	if val, ok = c.data[key]; ok {
		return val
	}

	c.data[key] = value

	return value
}

func main() {
	cm := NewConcurrentMap()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		val := cm.GetOrCreate("key1", "value1")

		fmt.Println("goroutine 1 got: ", val)
	}()

	go func() {
		defer wg.Done()

		val := cm.GetOrCreate("key1", "value2")

		fmt.Println("goroutine 2 got: ", val)
	}()

	wg.Wait()
}

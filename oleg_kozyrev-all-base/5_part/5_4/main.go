package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func (c *Cache) Store(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *Cache) Load(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.data[key]

	return value, ok
}

func main() {
	cache := &Cache{
		data: make(map[string]interface{}),
	}

	cache.Store("name", "Alice")
	cache.Store("age", 25)

	name, ok := cache.Load("name")
	if !ok {
		fmt.Println("Name not found")
	} else {
		nameStr, ok := name.(string)
		if !ok {
			fmt.Println("Name is not a string")
		} else {
			fmt.Println("Name:", nameStr)
		}
	}

	age, ok := cache.Load("age")
	if !ok {
		fmt.Println("Age not found")
	} else {
		ageInt, ok := age.(int)
		if !ok {
			fmt.Println("Age is not an int")
		} else {
			fmt.Println("Age:", ageInt)
		}
	}

	_, ok = cache.Load("height")
	if !ok {
		fmt.Println("Height not found")
	}

	width, ok := cache.Load("width")
	if !ok {
		fmt.Println("Width not found")
	} else {
		widthFl, ok := width.(float64)
		if !ok {
			fmt.Println("Width is not a float64")
		} else {
			fmt.Println("Width:", widthFl)
		}
	}
}

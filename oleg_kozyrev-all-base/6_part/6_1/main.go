package main

import (
	"context"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter atomic.Int64

func SimpleRequest(ctx context.Context) (int64, error) {
	start := time.Now()
	defer func() {
		log.Printf("Request time: %v\n", time.Since(start))
	}()

	ch := make(chan int64)

	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		counter.Add(1)
		ch <- counter.Load()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case counter := <-ch:
		return counter, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	value, err := SimpleRequest(ctx)
	if err != nil {
		log.Printf("Request failed: %v\n", err)
	}

	log.Printf("Counter value: %d\n", value)
}

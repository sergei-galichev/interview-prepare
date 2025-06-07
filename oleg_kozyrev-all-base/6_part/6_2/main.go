package main

import (
	"context"
	"fmt"
	"time"
)

const (
	defaultTimeout = 1 * time.Second
)

func getDiscount() float64 {
	time.Sleep(2 * time.Second)
	return 12.0
}

func getDiscountWithTimeout(ctx context.Context) (float64, error) {
	ch := make(chan float64)

	go func() {
		res := getDiscount()
		ch <- res
		close(ch)
	}()

	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
	}

	to, _ := ctx.Deadline()

	fmt.Printf("deadline: %v\n", time.Since(to))
	//fmt.Printf("deadline: %v\n", to)

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case result := <-ch:
		return result, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//ctx := context.Background()

	result, err := getDiscountWithTimeout(ctx)
	if err != nil {
		fmt.Printf("Something wrong: %v", err)
		return
	}

	fmt.Printf("Your discount: %v", result)
}

package main

import (
	"circuit_breaker"
	"context"
	"log/slog"
	"time"
)

func main() {
	ctx := context.Background()

	timer := time.NewTimer(time.Millisecond * 750)

	service := NewService()

	circuit := circuit_breaker.Breaker(service.CreateWork, 3)

	for i := 0; i < 35; i++ {
		work, err := circuit(ctx)
		if err != nil {
			slog.Error("create work failed", slog.String("error", err.Error()))
		} else {
			slog.Info("work created", slog.String("work", work))
		}

		<-timer.C

		timer.Reset(time.Millisecond * 750)
	}
}

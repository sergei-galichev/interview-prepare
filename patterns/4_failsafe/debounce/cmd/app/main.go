package main

import (
	"context"
	"debounce"
	"log/slog"
	"time"
)

func main() {
	ctx := context.Background()

	timer := time.NewTimer(time.Second * 1)

	service := NewService()

	df := debounce.FailsafeDebounceFirst(service.CreateWork, time.Second*2)

	for i := 0; i < 10; i++ {
		work, err := df(ctx)
		if err != nil {
			slog.Error("create work failed", slog.String("error", err.Error()))
		} else {
			slog.Info("work created", slog.String("work", work))
		}

		<-timer.C

		timer.Reset(time.Second * 3)
	}
}

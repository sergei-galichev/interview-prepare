package main

import (
	"context"
	"log/slog"
	"throttle"
	"time"
)

func main() {
	ctx := context.Background()

	timer := time.NewTimer(time.Millisecond * 100)

	service := NewService()

	th := throttle.Throttle(service.CreateWork, 50, 50, time.Second*10)

	for i := 0; i < 500; i++ {
		work, err := th(ctx)
		if err != nil {
			slog.Error("failed to create work", slog.String("error", err.Error()))
		} else {
			slog.Info("work created", slog.String("work", work), slog.Int("id", i))
		}
		<-timer.C

		timer.Reset(time.Millisecond * 100)
	}

}

package main

import (
	"context"
	"log/slog"
	"os"
	"retry"
	"time"
)

func main() {
	ctx := context.Background()

	service := NewService()

	r := retry.Retry(ctx, service.CreateWork, 5, time.Second*2)

	work, err := r(ctx)
	if err != nil {
		slog.Error("Retry failed", slog.String("error", err.Error()))

		os.Exit(1)
	} else {
		slog.Info("work created", slog.String("work", work))
	}
}

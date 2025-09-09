package retry

import (
	"context"
	"log/slog"
	"time"
)

type Effector func(ctx context.Context) (string, error)

func Retry(ctx context.Context, effector Effector, retries int, delay time.Duration) Effector {
	var threshold time.Duration

	return func(ctx context.Context) (string, error) {
		threshold = delay

		for r := 0; ; r++ {
			response, err := effector(ctx)
			if err == nil || r >= retries {
				return response, err
			}

			slog.Warn("Attempt failed", slog.Int("attempt", r+1), slog.String("delay", threshold.String()))

			select {
			case <-time.After(threshold):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}

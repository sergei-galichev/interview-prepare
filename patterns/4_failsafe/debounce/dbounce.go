package debounce

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

type Circuit func(ctx context.Context) (string, error)

func FailsafeDebounceFirst(c Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var mu sync.Mutex

	return func(ctx context.Context) (string, error) {
		mu.Lock()
		defer func() {
			threshold = time.Now().Add(d)

			slog.Info("threshold set", slog.String("value", threshold.String()))

			mu.Unlock()
		}()

		if time.Now().Before(threshold) {
			slog.Info("threshold not reached")

			return result, err
		}

		slog.Info("threshold reached")

		result, err = c(ctx)

		return result, err
	}
}

func FailsafeDebounceLast(c Circuit, d time.Duration) Circuit {
	var threshold time.Time = time.Now()
	var ticker *time.Ticker
	var result string
	var err error
	var once sync.Once
	var mu sync.Mutex

	return func(ctx context.Context) (string, error) {
		mu.Lock()
		defer mu.Unlock()

		threshold = time.Now().Add(d)

		once.Do(
			func() {
				ticker = time.NewTicker(time.Millisecond * 100)

				go func() {
					defer func() {
						mu.Lock()

						ticker.Stop()

						once = sync.Once{}

						mu.Unlock()
					}()

					for {
						select {
						case <-ticker.C:
							mu.Lock()

							if time.Now().After(threshold) {
								result, err = c(ctx)

								mu.Unlock()

								return
							}

							mu.Unlock()
						case <-ctx.Done():
							mu.Lock()

							result, err = "", ctx.Err()

							mu.Unlock()
							return
						}
					}
				}()
			},
		)

		return result, err
	}
}

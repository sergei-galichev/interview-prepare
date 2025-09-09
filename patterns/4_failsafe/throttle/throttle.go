package throttle

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Effector func(ctx context.Context) (string, error)

func Throttle(effector Effector, maxTokens, refill uint, d time.Duration) Effector {
	var tokens = maxTokens
	var once sync.Once

	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		once.Do(
			func() {
				ticker := time.NewTicker(d)

				go func() {
					defer ticker.Stop()

					for {
						select {
						case <-ctx.Done():
							return
						case <-ticker.C:
							t := tokens + refill

							if t > maxTokens {
								t = maxTokens
							}

							tokens = t
						}
					}
				}()
			},
		)

		if tokens <= 0 {
			return "", errors.New("rate limit exceeded")
		}

		tokens--

		return effector(ctx)
	}
}

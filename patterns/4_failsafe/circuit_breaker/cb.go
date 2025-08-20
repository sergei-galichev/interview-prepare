package circuit_breaker

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"
)

type Circuit func(ctx context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock()

		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 {
			dur := time.Second * 2 << d

			shouldRetryAt := lastAttempt.Add(dur)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()

				slog.Error("circuit is open", slog.Duration("duration", dur))

				return "", errors.New("service is unavailable")
			}
		}

		m.RUnlock()

		res, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++

			return res, err
		}

		consecutiveFailures = 0

		return res, nil
	}
}

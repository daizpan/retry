package retry

import (
	"context"
	"time"
)

func Do(n uint, interval time.Duration, fn func() error) (err error) {
	return RetryWithContext(context.Background(), n, interval, fn)
}

func RetryWithContext(ctx context.Context, n uint, interval time.Duration, fn func() error) (err error) {
	var count int
	for {
		count++
		err = fn()
		if err == nil || uint(count) >= n {
			return
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}
	}
}

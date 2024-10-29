package utils

import (
	"context"
	"time"
)

// created this as a helper function to call a gateway function with retry logic
func CallGatewayWithRetry(ctx context.Context, retryCount int, timeout time.Duration, gatewayFunc func(ctx context.Context) error) error {
	var lastErr error
	for i := 0; i < retryCount; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			// attempt to call the gateway function with the timeout context
			err := gatewayFunc(timeoutCtx)
			if err == nil {
				return nil
			}

			lastErr = err
			time.Sleep(100 * time.Millisecond) // small wait before doing a retry
		}
	}
	return lastErr
}

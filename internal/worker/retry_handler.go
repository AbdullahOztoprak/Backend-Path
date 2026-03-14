//go:build worker
// +build worker

package worker

import (
	"log"
	"time"
)

type RetryHandler struct {
	maxRetries int
	delay      time.Duration
}

func NewRetryHandler(maxRetries int, delay time.Duration) *RetryHandler {
	return &RetryHandler{
		maxRetries: maxRetries,
		delay:      delay,
	}
}

func (r *RetryHandler) ExecuteWithRetry(operation func() error) error {
	var err error
	for i := 0; i < r.maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		log.Printf("Operation failed: %v. Retrying in %v...", err, r.delay)
		time.Sleep(r.delay)
	}
	return err
}
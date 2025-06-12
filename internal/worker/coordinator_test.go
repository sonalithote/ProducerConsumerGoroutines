package worker

import (
	"context"
	"testing"
	"time"
)

func TestRunProducerConsumer_NormalCompletion(t *testing.T) {
	ctx := context.Background()
	bufferSize := 5
	totalItems := 8
	numProducers := 2
	numConsumers := 2

	RunProducerConsumer(ctx, bufferSize, totalItems, numProducers, numConsumers)
	// If no deadlock or panic, test passes
}

func TestRunProducerConsumer_Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	bufferSize := 5
	totalItems := 100
	numProducers := 2
	numConsumers := 2

	go func() {
		time.Sleep(200 * time.Millisecond)
		cancel()
	}()

	RunProducerConsumer(ctx, bufferSize, totalItems, numProducers, numConsumers)
	// Should exit gracefully on cancellation
}

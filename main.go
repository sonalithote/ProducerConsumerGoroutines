package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sonalithote/producer-consumer/internal/worker"
)

func main() {
	bufferSize := 5
	totalItems := 20
	numProducers := 3
	numConsumers := 2

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS interrupt for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		cancel()
	}()

	// Optional: Cancel after a timeout (for demo)
	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	worker.RunProducerConsumer(ctx, bufferSize, totalItems, numProducers, numConsumers)
}

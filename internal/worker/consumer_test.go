package worker

import (
	"context"
	"testing"
	"time"
)

func TestConsumer_NormalCompletion(t *testing.T) {
	ch := make(chan int, 5)
	done := make(chan bool)
	ctx := context.Background()

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	go Consumer(ctx, 1, ch, done)
	select {
	case <-done:
		// Success
	case <-time.After(2 * time.Second):
		t.Error("Consumer did not finish in time")
	}
}

func TestConsumer_Cancellation(t *testing.T) {
	ch := make(chan int, 5)
	done := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()

	go Consumer(ctx, 1, ch, done)

	time.Sleep(120 * time.Millisecond)
	cancel() // Cancel context

	select {
	case <-done:
		// Success
	case <-time.After(2 * time.Second):
		t.Error("Consumer did not exit upon cancellation")
	}
}

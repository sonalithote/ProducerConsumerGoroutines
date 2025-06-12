package worker

import (
	"context"
	"testing"
	"time"
)

func TestProducer_NormalCompletion(t *testing.T) {
	ch := make(chan int, 5)
	ctx := context.Background()

	go func() {
		Producer(ctx, 1, ch, 3)
		close(ch)
	}()

	var produced []int
	for item := range ch {
		produced = append(produced, item)
	}

	if len(produced) != 3 {
		t.Errorf("Expected 3 items, got %d", len(produced))
	}
}

func TestProducer_Cancellation(t *testing.T) {
	ch := make(chan int, 5)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		Producer(ctx, 1, ch, 10)
		close(ch)
	}()

	time.Sleep(150 * time.Millisecond) // Let it produce a few items
	cancel()                           // Cancel context

	count := 0
	for range ch {
		count++
	}
	if count == 0 {
		t.Error("Producer did not produce any items before cancellation")
	}
}

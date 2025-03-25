package worker

import (
	"testing"
)

func TestProducer(t *testing.T) {
	bufferSize := 5
	ch := make(chan int, bufferSize)
	itemsPerProducer := 10

	go func() {
		defer close(ch)
		Producer(1, ch, itemsPerProducer)
	}()

	producedItems := []int{}
	for item := range ch {
		producedItems = append(producedItems, item)
	}

	if len(producedItems) != itemsPerProducer {
		t.Errorf("Expected %d items, but got %d", itemsPerProducer, len(producedItems))
	}
}

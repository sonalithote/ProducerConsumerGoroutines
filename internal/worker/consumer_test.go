package worker

import (
	"testing"
)

func TestConsumer(t *testing.T) {
	bufferSize := 5
	ch := make(chan int, bufferSize)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go Consumer(1, ch, done)

	if <-done != true {
		t.Error("Consumer did not signal completion correctly")
	}
}

package worker

import (
	"fmt"
	"time"
)

// Consumer processes items from the channel.
func Consumer(id int, ch <-chan int, done chan<- bool) {
	for item := range ch {
		fmt.Printf("Consumer %d: Consumed item %d\n", id, item)
		time.Sleep(200 * time.Millisecond) // Simulate processing delay
	}
	done <- true // Signal that consumer is finished
}

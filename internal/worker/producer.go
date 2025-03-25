package worker

import (
	"fmt"
	"time"
)

// Producer generates items and sends them to the channel.
func Producer(id int, ch chan<- int, items int) {
	for i := 0; i < items; i++ {
		item := id*100 + i // Unique item ID based on producer ID
		fmt.Printf("Producer %d: Produced item %d\n", id, item)
		ch <- item
		time.Sleep(100 * time.Millisecond) // Simulate production delay
	}
}

package worker

import (
	"context"
	"fmt"
	"time"
)

// Producer generates items and sends them to the channel.
func Producer(ctx context.Context, id int, ch chan<- int, items int) {
	for i := 0; i < items; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Producer %d: cancelled\n", id)
			return
		case ch <- id*100 + i:
			fmt.Printf("Producer %d: produced %d\n", id, id*100+i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

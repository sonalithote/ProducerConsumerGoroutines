package worker

import (
	"context"
	"fmt"
	"time"
)

// Consumer processes items from the channel.
func Consumer(ctx context.Context, id int, ch <-chan int, done chan<- bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Consumer %d: cancelled\n", id)
			done <- true
			return
		case item, ok := <-ch:
			if !ok {
				fmt.Printf("Consumer %d: channel closed, exiting\n", id)
				done <- true
				return
			}
			fmt.Printf("Consumer %d: consumed %d\n", id, item)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

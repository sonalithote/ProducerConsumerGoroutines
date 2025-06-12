package worker

import (
	"context"
	"fmt"
	"sync"
)

// RunProducerConsumer starts multiple producers and consumers using goroutines.
func RunProducerConsumer(ctx context.Context, bufferSize int, totalItems int, numProducers int, numConsumers int) {
	ch := make(chan int, bufferSize)
	done := make(chan bool)
	var wg sync.WaitGroup

	itemsPerProducer := totalItems / numProducers

	// Start producers
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Producer(ctx, id, ch, itemsPerProducer)
		}(i)
	}

	// Close channel when all producers are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Start consumers
	for i := 1; i <= numConsumers; i++ {
		go Consumer(ctx, i, ch, done)
	}

	// Wait for all consumers to finish
	for i := 1; i <= numConsumers; i++ {
		<-done
	}

	fmt.Println("All items processed or cancelled.")
}

package worker

import (
	"fmt"
	"sync"
)

// RunProducerConsumer starts multiple producers and consumers using goroutines.
func RunProducerConsumer(bufferSize int, totalItems int, numProducers int, numConsumers int) {
	ch := make(chan int, bufferSize)
	done := make(chan bool)
	var wg sync.WaitGroup

	itemsPerProducer := totalItems / numProducers

	// Start producers
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Producer(id, ch, itemsPerProducer)
		}(i)
	}

	// Start a goroutine to close the channel after all producers finish
	go func() {
		wg.Wait() // Wait for all producers to finish
		close(ch) // Close the channel so consumers stop
	}()

	// Start consumers
	for i := 1; i <= numConsumers; i++ {
		go Consumer(i, ch, done)
	}

	// Wait for all consumers to finish
	for i := 1; i <= numConsumers; i++ {
		<-done
	}

	fmt.Println("All items processed.")
}

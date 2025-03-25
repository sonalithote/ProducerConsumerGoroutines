package main

import (
	"github.com/sonalithote/producer-consumer/internal/worker"
)

func main() {
	bufferSize := 5
	itemsToProduce := 20
	numProducers := 3
	numConsumers := 2

	worker.RunProducerConsumer(bufferSize, itemsToProduce, numProducers, numConsumers)
}

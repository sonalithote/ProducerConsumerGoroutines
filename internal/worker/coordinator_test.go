package worker

import (
	"testing"
)

func TestRunProducerConsumer(t *testing.T) {
	bufferSize := 5
	totalItems := 20
	numProducers := 3
	numConsumers := 2

	RunProducerConsumer(bufferSize, totalItems, numProducers, numConsumers)

	t.Log("Integration test passed if no deadlocks or errors occurred.")
}

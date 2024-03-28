package batch

import (
	. "github.com/Shibbaz/GO-EventBus/pkg/events"
)

func (batch *Batch) Publish(data *[]Event, batchSize int, dispatcher *Dispatcher) {
	batches := batch.Push(data, batchSize)
	for _, events := range batches {
		for _, event := range events {
			event.Exec(dispatcher)
		}
	}
}

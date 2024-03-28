package batch

import (
	. "github.com/Shibbaz/GO-EventBus/pkg/events"
)

func (batch *Batch) Publish(events *[]Event, batchSize int, dispatcher *Dispatcher) {
	for _, batch := range batch.Push(*events, batchSize) {
		for _, event := range batch {
			event.Exec(dispatcher)
		}
	}
}

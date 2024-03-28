package batch

import (
	. "github.com/Shibbaz/GO-EventBus/pkg/events"
)

func (batch *Batch) Publish(events *[]Event, batchSize int, dispatcher *Dispatcher) {
	// batch [][]Event where each element is []Event => event, read below
	for _, batch := range batch.Push(*events, batchSize) {
		// event = []Event of max size batchSize
		for _, event := range batch {
			event.Exec(dispatcher)
		}
	}
}

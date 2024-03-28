package batch

import (
	. "github.com/Shibbaz/GO-EventBus/pkg/events"
)

func (batch Batch) Push(events []Event, batchSize int) Batch {
	store := NewStore()
	// add go routine to fasten the process
	for _, event := range events {
		store.Subscribe(event)
		if len(store.Events) == batchSize {
			batch = append(batch, store.Events)
			store.Reset()
			continue
		}
	}
	batch = append(batch, store.Events)
	return batch
}

package batch

import (
	. "github.com/Shibbaz/GO-EventBus/internal/events"
)

func (batch Batch) Push(events []Event, batchSize int) Batch {
	store := NewStore()
	for loop := true; loop; {
		for _, event := range events {
			store.Subscribe(event)
			if len(store.Events) == batchSize {
				batch = append(batch, store.Events)
				store.Reset()
				continue
			}
		}
		batch = append(batch, store.Events)

		loop = false
	}

	return batch
}

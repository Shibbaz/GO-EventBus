package batch

import (
	. "events"
)

func (batch Batch) Push(events []Event, batchSize int) Batch {
	batchedEvents := []Event{}
	for loop := true; loop; {
		for _, event := range events {
			batchedEvents = append(batchedEvents, event)
			if len(batchedEvents) == batchSize {
				batch = append(batch, batchedEvents)
				batchedEvents = []Event{}
				continue
			}
		}
		batch = append(batch, batchedEvents)

		loop = false
	}

	return batch
}

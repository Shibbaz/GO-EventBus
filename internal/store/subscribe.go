package store

import (
	. "events"
)

func (store *EventsStore) Subscribe(event *Event) {
	store.Stream = append(store.Stream, *event)
}

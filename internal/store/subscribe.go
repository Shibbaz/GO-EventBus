package store

import(
	. "events"
)

func (store *EventsStore) Subscribe(event *Event) {
	store.Events = append(store.Events, *event);
}
package store

import(
	. "events"
)

func NewEventsStore(dispatcher EventDispatcher) *EventsStore{
	return &EventsStore{
		Events: []Event{},
		Dispatcher: dispatcher,
	}
}
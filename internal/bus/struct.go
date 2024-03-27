package bus

import (
	. "dispatcher"
	. "events"
)

type Bus struct {
	Channels   map[int](chan map[string]any)
	Events     EventList
	Dispatcher Dispatcher
}

func NewBus(dispatcher *Dispatcher) *Bus {
	channel := make(map[int](chan map[string]any))

	return &Bus{
		Channels:   channel,
		Events:     *NewEventList(),
		Dispatcher: *dispatcher,
	}
}

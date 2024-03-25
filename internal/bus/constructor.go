package bus

import (
	. "dispatcher"
	. "events"
)

func NewBus(dispatcher *Dispatcher) *Bus {
	channel := make(map[int](chan map[string]any))

	return &Bus{
		Channels:   channel,
		Events:     []Event{},
		Dispatcher: *dispatcher,
	}
}

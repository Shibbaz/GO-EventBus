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

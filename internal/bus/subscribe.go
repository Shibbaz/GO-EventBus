package bus

import (
	. "events"
)

func (bus *Bus) Subscribe(event *Event) {
	bus.Events = append(bus.Events, *event)

}

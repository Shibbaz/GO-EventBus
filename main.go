package main

import (
	. "bus"
	. "dispatcher"
	. "events"
	. "examples"
	"reflect"
)

func main() {
	dispatcher := Dispatcher{
		reflect.TypeOf(Projection{}): Example,
	}
	bus := NewBus(&dispatcher)

	event := NewEvent(Projection{}, EventArgs{1: "1"})

	bus.Subscribe(event)
	bus.Compose()
}

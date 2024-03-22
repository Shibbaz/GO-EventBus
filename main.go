package main

import (
	. "events"
	. "examples"
	"reflect"
	. "store"
	. "types"
)

func main() {
	EventDispatcher := NewEventDispatcher(
		map[reflect.Type]func(EventArgs) (Status, error){
			reflect.TypeOf(Projection{}): Example,
		},
	)
	eventStore := EventsStore{
		Dispatcher: *EventDispatcher,
	}
	eventStore.Subscribe(&Event{
		Projection: *NewProjection(),
		Args: EventArgs{
			"1": 1,
			"2": 2,
		},
	})
	eventStore.Publish()
}

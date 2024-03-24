package store

import (
	. "events"
	"reflect"
	. "types"
)

type EventsStore struct {
	Stream     []Event
	Dispatcher EventDispatcher
}

type EventDispatcher struct {
	EventFunctions map[reflect.Type]func(EventArgs) (Status, error)
}

package store

import(
	. "types"
	"reflect"
	. "events"
)

type EventsStore struct{
	Events []Event
	Dispatcher EventDispatcher
}

type EventDispatcher struct{
	EventFunctions map[reflect.Type] func(EventArgs)(Status, error)
}
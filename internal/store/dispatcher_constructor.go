package store

import(
	"reflect"
	. "types"
)
func NewEventDispatcher(setup map[reflect.Type] func(EventArgs)(Status, error)) *EventDispatcher{
	return &EventDispatcher{
		EventFunctions: setup,
	}
}
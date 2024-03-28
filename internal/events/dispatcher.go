package events

import (
	"reflect"
)

type Dispatcher map[reflect.Type]func(EventArgs) error

func (dispatcher *Dispatcher) Get(projection any) func(EventArgs) error {
	return (*dispatcher)[reflect.TypeOf(projection)]
}

package dispatcher

import (
	. "events"
	"reflect"
)

func (dispatcher *Dispatcher) Get(projection any) func(EventArgs) error {
	return (*dispatcher)[reflect.TypeOf(projection)]
}

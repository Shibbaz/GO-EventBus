package dispatcher

import (
	. "events"
	"reflect"
)

func (dispatcher *Dispatcher) Get(projection any) func(EventArgs) (Data, error) {
	return (*dispatcher)[reflect.TypeOf(projection)]
}

package dispatcher

import (
	. "events"
	"reflect"
)

type Data struct {
	Metadata string
}
type Dispatcher map[reflect.Type]func(EventArgs) (Data, error)

func (dispatcher *Dispatcher) Get(projection any) func(EventArgs) (Data, error) {
	return (*dispatcher)[reflect.TypeOf(projection)]
}

package dispatcher

import "reflect"

type Data struct {
	Metadata string
}
type Dispatcher map[reflect.Type]func(...any) (Data, error)

func (dispatcher *Dispatcher) Get(projection any) func(...any) (Data, error) {
	return (*dispatcher)[reflect.TypeOf(projection)]
}

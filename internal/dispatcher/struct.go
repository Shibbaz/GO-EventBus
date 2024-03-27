package dispatcher

import (
	. "events"
	"reflect"
)

type Data struct {
	Metadata string
}
type Dispatcher map[reflect.Type]func(EventArgs) (Data, error)

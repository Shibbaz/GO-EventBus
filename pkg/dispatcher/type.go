package dispatcher

import (
	. "events"
	"reflect"
)

type Dispatcher map[reflect.Type]func(EventArgs) error

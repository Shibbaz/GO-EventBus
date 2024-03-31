package GOEventBus

import "reflect"

type Dispatcher map[reflect.Type]func(EventArgs) error

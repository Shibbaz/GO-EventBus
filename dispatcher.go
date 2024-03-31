package GOEventBus

import "reflect"

// Dispatcher contains projections as keys and func(EventArgs) error as value
// Store recognizes which action to take depending on events' projection
type Dispatcher map[reflect.Type]func(EventArgs) error

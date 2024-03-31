package GOEventBus

type EventArgs map[string]any

// Event provides args for event funcs, just by EventArgs{"key1": value, "key2":value}
// Projection is empty struct, dispatcher links projection to func to execute by publishing event
type Event struct {
	projection any
	args       EventArgs
}

func NewEvent(projection any, args EventArgs) *Event {
	return &Event{
		projection: projection,
		args:       args,
	}
}

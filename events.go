package GOEventBus

type EventArgs map[string]any
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

package events

type EventArgs map[int]any
type Event struct {
	Projection any
	Args       EventArgs
}

func NewEvent(projection any, args EventArgs) Event {
	return Event{
		Projection: projection,
		Args:       args,
	}
}

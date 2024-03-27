package events

type EventArgs []any

type Event struct {
	Projection any
	EventArgs  EventArgs
	Status     bool
}

type EventList []Event

func NewEvent(projection any, args EventArgs) *Event {
	return &Event{
		Projection: projection,
		EventArgs:  args,
		Status:     false,
	}
}

func NewEventList() *EventList {
	return &EventList{}
}

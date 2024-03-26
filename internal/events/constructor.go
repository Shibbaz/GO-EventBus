package events

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

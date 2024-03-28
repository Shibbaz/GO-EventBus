package events

func NewEvent(args EventArgs, projection any) Event {
	return Event{
		Args:       args,
		Projection: projection,
	}
}

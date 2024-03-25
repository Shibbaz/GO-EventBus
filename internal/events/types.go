package events

type EventArgs []any

type Event struct {
	Projection any
	EventArgs  EventArgs
}

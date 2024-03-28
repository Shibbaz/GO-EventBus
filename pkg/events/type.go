package events

type EventArgs map[string]any
type Event struct {
	Args       EventArgs
	Projection any
}

package events

func (event *Event) Exec(dispatcher *Dispatcher) {
	dispatcher.Get(event.Projection)(event.Args)
}

package events

type Store struct {
	Events []Event
}

func NewStore() *Store {
	return &Store{
		Events: []Event{},
	}
}

func (store *Store) Publish(dispatcher *Dispatcher) {
	for _, event := range store.Events {
		event.Exec(dispatcher)
	}
}

func (store *Store) Subscribe(event Event) {
	store.Events = append(store.Events, event)
}

func (store *Store) Reset() {
	store.Events = []Event{}
}

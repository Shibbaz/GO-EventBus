package store

func (store *EventsStore) Publish(){
	for _, event := range store.Events {
		store.GetFunc(event)(event.Args);
	}
}
package store

import (
	"fmt"
	"reflect"
)

func (store *EventsStore) Publish() {
	for _, event := range store.Events {
		fn := store.GetFunc(event.Projection)
		typeOf := reflect.TypeOf(event.Projection).Name()
		fmt.Printf("Event type of %s got results", typeOf)
		fn(event.Args)
	}
}

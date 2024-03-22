package store

import(
	"reflect"
	. "types"
)

func (store *EventsStore) GetFunc(projection any) func(EventArgs)(Status, error){
	typeOfProjection := reflect.TypeOf(projection);
	return store.Dispatcher.EventFunctions[typeOfProjection]
}
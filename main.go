package main

import(
	"reflect"
	"fmt"
	. "store"
	. "examples"
	. "types"
)


func main(){
	EventDispatcher := NewEventDispatcher(
		map[reflect.Type] func(EventArgs)(Status, error){
			reflect.TypeOf(Projection{}): Example,
		},
	);
	eventStore := EventsStore{
		Dispatcher: *EventDispatcher,
	}
	method := eventStore.GetFunc(Projection{});
	args := EventArgs{
		"xd": 1,
	}
	fmt.Println(method(args))
	
}
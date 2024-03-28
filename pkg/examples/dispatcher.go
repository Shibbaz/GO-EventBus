package examples

import (
	. "dispatcher"
	. "events"
	"fmt"
	"reflect"
)

func SellingHouse(args EventArgs) error {
	fmt.Printf("House %d was sold for %v aud\n", args["id"], args["price"])
	return nil
}

var EventsDispatcher = Dispatcher{
	reflect.TypeOf(HouseWasSold{}): SellingHouse,
}

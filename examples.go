package GOEventBus

import (
	"fmt"
)

func Example(args *map[string]any) (map[string]any, error) {
	fmt.Println(args)

	return *args, nil
}

var ExampleDispatcher = Dispatcher{
	"HouseWasSold": Example,
}

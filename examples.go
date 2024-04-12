package GOEventBus

import "fmt"

func Example(args map[string]any) {
	fmt.Println(args)
}

var ExampleDispatcher = Dispatcher{
	"HouseWasSold": Example,
}

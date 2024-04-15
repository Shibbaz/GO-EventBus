package GOEventBus

import (
	"fmt"
)

func Example(args map[string]any) ([]byte, error) {
	fmt.Println(args)

	var data []byte = (&Serializer{}).Serialize(args)

	return data, nil
}

var ExampleDispatcher = Dispatcher{
	"HouseWasSold": Example,
}

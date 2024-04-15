package GOEventBus

import (
	"fmt"
	"unsafe"
)

func Example(args map[string]any) ([]byte, error) {
	fmt.Println(args)

	var data []byte = *(*[]byte)(unsafe.Pointer(&args))

	return data, nil
}

var ExampleDispatcher = Dispatcher{
	"HouseWasSold": Example,
}

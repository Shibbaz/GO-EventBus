package GOEventBus

import (
	"fmt"
	"unsafe"
)

func Example(args map[string]any) []byte {
	fmt.Println(args)

	var data []byte = *(*[]byte)(unsafe.Pointer(&args))

	return data
}

var ExampleDispatcher = Dispatcher{
	"HouseWasSold": Example,
}

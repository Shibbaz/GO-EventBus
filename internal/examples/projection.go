package examples

import (
	. "dispatcher"
)

type Projection struct{}

func Example(...any) (Data, error) {
	return Data{Metadata: "Hello"}, nil
}

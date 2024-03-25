package examples

import (
	. "dispatcher"
	"fmt"
)

type Projection struct{}

func Example(args ...any) (Data, error) {
	return Data{Metadata: fmt.Sprintf("%v", args)}, nil
}

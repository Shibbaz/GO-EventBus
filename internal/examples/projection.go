package examples

import (
	. "dispatcher"
	. "events"
	"fmt"
)

type Projection struct{}

func Example(args EventArgs) (Data, error) {
	return Data{Metadata: fmt.Sprintf("%v", args)}, nil
}

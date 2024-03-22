package examples

import (
	"fmt"
	. "types"
)

func Example(a EventArgs) (Status, error) {
	status := *NewStatus(200, "Succesful")
	fmt.Printf(" %v %v\n", a, status)
	return status, nil
}

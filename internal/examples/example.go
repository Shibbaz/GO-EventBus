package examples

import(
	"fmt"
	. "types"
)

func Example(a EventArgs)(Status, error){
	fmt.Println(a)
	return Status{
		Code: 200,
		Message: "Successful",
	},nil
}
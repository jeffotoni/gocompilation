// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"fmt"
)

type MyStruct struct {
	Msg interface{}
}

func main() {
	b := MyStruct{}
	// string
	b.Msg = "5"
	fmt.Printf("%#v %T \n", b.Msg, b.Msg) // Output: "5" string

	// int
	b.Msg = 5
	fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int

	// map
	b.Msg = map[string]string{"population": "500000", "language": "sueco"}
	fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int
}

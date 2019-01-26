package main

import "fmt"

func main() {
	var a int = 100  /* actual variable declaration */
	var pa *int      /* pointer variable declaration */
	var pointer *int /* pointer variable declaration */

	pa = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", pa)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *pa)

	/* succeeds if p is not nil */
	if pa != nil {
		fmt.Println("success is not nil")
	}

	/* succeeds if p is null */
	if (pointer) == nil {
		fmt.Println("success pointer is nil")
	}
}

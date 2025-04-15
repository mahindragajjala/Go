package main

import (
	"fmt"
	referencedereference "pointer/reference_dereference"
)

func main() {
	var data int
	var data_pointer = &data
	fmt.Println(*data_pointer)

	referencedereference.Reference()
}

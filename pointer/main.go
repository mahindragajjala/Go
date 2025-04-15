package main

import (
	"fmt"
)

// func main() {
// 	var data int
// 	var data_pointer = &data
// 	fmt.Println(*data_pointer)

//		referencedereference.Reference()
//	}
func main() {
	var a int = 10
	var p *int = &a // Reference: store address of 'a' into pointer 'p'
	b := &a

	//cannot use &a (value of type *int) as int value in variable declarationcompilerIncompatibleAssign

	fmt.Println("Address of a:", p)
	fmt.Println(b)
}

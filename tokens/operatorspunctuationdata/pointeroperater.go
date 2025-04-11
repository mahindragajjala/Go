package operatorspunctuationdata

import "fmt"

func Pointer_operator() {
	//Dereferencing a pointer
	var x int = 10
	var p *int = &x // p is a pointer to x
	fmt.Println(*p) // Outputs: 10

	//Declaring a pointer
	var Declare_pointer *int
	fmt.Printf("Declare_pointer---->:%T", Declare_pointer)

	//Multiplication

	var y int = 10
	var z int = x * y
	fmt.Println(z)
}

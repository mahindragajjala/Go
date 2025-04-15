package code

import "fmt"

func modify(num *int) {
	*num = 100
}

func PassByAddress() {
	x := 10
	fmt.Println("Before:", x) // Output: 10

	modify(&x) // Passing the address of x

	fmt.Println("After:", x) // Output: 100
}

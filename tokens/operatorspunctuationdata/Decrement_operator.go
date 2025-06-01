package operatorspunctuationdata

import "fmt"

func Decrement_Operator() {
	/*
		x-- is used to decrease the value of x by 1.
		It is equivalent to:
		x = x - 1
	*/
	x := 5
	x-- // Now x = 4
	fmt.Println(x)

	Z := 5

	var count int = 5

	for i := 0; i < Z; i++ {
		count--
	}
}

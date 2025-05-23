package operatorspunctuationdata

import "fmt"

//x = x &^ y
func AND_NOT_assignment_operator() {
	x := 0b1111 // 15 in decimal
	y := 0b0101 // 5 in decimal

	x &^= y
	fmt.Printf("%04b\n", x) // Output: 1010 (10 in decimal)

	/*
		x = 1111
		y = 0101
		----------
		x &^= y → 1010

	*/
}

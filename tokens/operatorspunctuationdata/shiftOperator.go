package operatorspunctuationdata

import "fmt"

func ShiftOperator() {
	//<<
	/*

		the << operator is the bitwise left shift operator.
		It shifts the bits to the left by a certain number of positions.
		Each left shift multiplies the number by 2ⁿ (power of 2).

		result := a << n

		a is the number you're shifting.
		n is how many positions to shift left.
		Using the << (left shift) operator in Go actually changes the binary
		representation of the number —
		by shifting its bits to the left, and filling in zeros on the right.


	*/
	a := 3
	b := a << 5
	fmt.Println(b) // Output: 6

	/*
		>> - bitwise right shift operator.

		It shifts all bits to the right by a certain number of positions.
		Each right shift divides the number by 2ⁿ (ignoring remainders).


	*/
	c := 8      // Binary: 00001000
	d := c >> 1 // Shift right by 1 → 00000100

	fmt.Println(d) // Output: 4

}

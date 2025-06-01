package operatorspunctuationdata

import "fmt"

func Bitwise_right_shift_assignment_operator() {

	var x uint = 16 // binary: 0001 0000
	x >>= 2         // shift right by 2 bits -> binary: 0000 0100 (which is 4)
	// x = x >> 2
	fmt.Println(x) // Output: 4

}

/*
Use case:

Bit shifting is common when working with:

Low-level data manipulation

Protocols

Flags or masks

Cryptography

Compression algorithms

*/

package operatorspunctuationdata

import "fmt"

func Kar_etOperator() {
	//Bitwise NOT (Unary Operator)
	a := 5       // Binary: 00000101
	result := ^a // Flips to: 11111010 (in 2’s complement: -6)

	fmt.Println(result) // Output: -6

	//Bitwise XOR (Binary Operator)
	c := 6           // 110
	d := 3           // 011
	resultd := c ^ d // 101 (which is 5)

	fmt.Println(resultd) // Output: 5

}

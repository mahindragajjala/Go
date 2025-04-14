package operatorspunctuationdata

import "fmt"

func Bit_clear() {
	x := 0b1101 // 13 in decimal
	y := 0b0101 // 5 in decimal

	z := x &^ y
	fmt.Printf("%04b\n", z) // Output: 1000

}

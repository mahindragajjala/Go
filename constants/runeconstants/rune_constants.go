package runeconstants

import "fmt"

/*

rune constants represent individual Unicode code points (characters)
and are of type rune, which is just an alias for int32.

What is a Rune?

A rune in Go:

Represents a Unicode code point.

Is an alias for int32.

Declared using single quotes: 'a', '₹', '中'

*/
func Rune_constant() {
	var ch rune = 'A'      // ASCII character
	var symbol rune = '₹'  // Unicode symbol
	var chinese rune = '中' // Non-ASCII character

	fmt.Println(ch)        // 65 (ASCII code)
	fmt.Printf("%c\n", ch) // A (character)

	fmt.Println(symbol)        // 8377
	fmt.Printf("%c\n", symbol) // ₹

	fmt.Println(chinese)        // 20013
	fmt.Printf("%c\n", chinese) // 中
}

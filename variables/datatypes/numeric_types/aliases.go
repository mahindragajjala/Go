package numerictypes

import "fmt"

/*

Alias	ActualType	Meaning/Usage
byte	uint8		Represents a single byte of data (0–255). Mostly used for binary data or ASCII characters.
rune	int32		Represents a Unicode code point. Mostly used for characters beyond ASCII (UTF-8 compatible).

*/
func Alias() {
	var b byte = 'A'              // ASCII character
	fmt.Println("byte value:", b) // Output: 65

	var r rune = '界'               // Unicode character (Chinese)
	fmt.Println("rune value:", r)  // Output: 30028
	fmt.Printf("Unicode: %U\n", r) // Output: U+754C
}

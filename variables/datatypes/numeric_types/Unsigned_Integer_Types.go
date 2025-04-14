package numerictypes

import (
	"fmt"
	"unsafe"
)

/*

Unsigned integers can only hold non-negative values (0 and above).

Type	Size	Range
uint8	8 bits	0 to 255
uint16	16 bits	0 to 65535
uint32	32 bits	0 to 4,294,967,295
uint64	64 bits	0 to 18,446,744,073,709,551,615
These are useful when you know the value won’t be negative, like RGB values or indexing arrays.
uint
An unsigned integer type (only positive values)
Size is also 32 or 64 bits depending on the system
*/

func Unsigned_integer_type() {
	var b uint16 = 60000
	fmt.Println("uint16 value:", b)

	var c uint32 = 4294967295
	fmt.Println("uint32 value:", c)

	var d uint64 = 18446744073709551615
	fmt.Println("uint64 value:", d)

	var e byte = 'A' // ASCII value of 'A' is 65
	fmt.Println("byte value:", e)

	var f uint = 1000
	fmt.Println("uint value:", f)
	fmt.Println("Size of uint:", unsafe.Sizeof(f)) // in bytes
}

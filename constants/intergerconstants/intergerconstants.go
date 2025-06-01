package intergerconstants

import "fmt"

/*

Concept	                    Example

Decimal	                    100
Binary (Go 1.13+)	        0b1010
Octal (Go 1.13+)	        0o77
Hexadecimal	                0xFF
Typed constant	            const x int = 5
Untyped constant	        const y = 5

*/
//Typed and Untyped
/*
Untyped Integer Constant
An untyped integer constant:

Has a value, but no fixed type yet.

It can adapt to the type expected by the context.
*/
const x = 10 // untyped integer constant

var a int = x     // x becomes int
var b int64 = x   // x becomes int64
var c float64 = x // x becomes float64

//Typed Integer Constant
//A typed integer constant has a fixed type when declared:
const y int32 = 10 // typed constant

// Integer Constants - These are untyped until used with a specific type like int, int32, int64, uint, etc.

const ab = 10     // untyped integer constant
const bc int = 20 // typed integer constant
const cd = -5

//numeric contants - You can also write integer constants in different number systems

const dec = 42       // decimal
const bin = 0b101010 // binary (Go 1.13+)
const oct = 0o52     // octal  (Go 1.13+)
const hex = 0x2A     // hexadecimal

func Interconstants() {
	fmt.Println(x, a, b, c, y, ab, bc, cd, dec, bin, oct, hex)
}

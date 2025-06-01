package operatorspunctuationdata

import "fmt"

/*
Operator	     Meaning	           Equivalent To

&=			Bitwise AND assignment	   x = x & y
`	                 =`	               Bitwise OR assignment
^=	        Bitwise XOR assignment	   x = x ^ y
<<=	        Left shift assignment	   x = x << y
>>=	        Right shift assignment	   x = x >> y

*/

//Flags (e.g., feature toggles)
const (
	FLAG_READ  = 1 << 0 // 0001
	FLAG_WRITE = 1 << 1 // 0010
	FLAG_EXEC  = 1 << 2 // 0100
)

func Flags() {
	var perms uint = 0

	// Enable read and write
	perms |= FLAG_READ
	perms |= FLAG_WRITE

	fmt.Printf("Permissions: %03b\n", perms) // Output: 011

	// Disable write
	perms &^= FLAG_WRITE // Bit clear (AND NOT)

	fmt.Printf("Permissions after removal: %03b\n", perms) // Output: 001
}

//Toggling bits (using XOR)
func Toggling_bits() {
	x := 0b1010
	y := 0b1100

	x ^= y // XOR: 1010 ^ 1100 = 0110
	//x = x ^ y
	fmt.Printf("%04b\n", x) // Output: 0110

	/*
	   Used in:

	   Cryptography

	   Checksums

	   Toggling bits
	*/
}

//Shifting bits (<<= and >>=)

func Shifting_bits() {
	x := uint(4) // 00000100

	x <<= 2        // Shift left by 2: 00010000 => 16
	fmt.Println(x) // 16

	x >>= 3        // Shift right by 3: 00000010 => 2
	fmt.Println(x) // 2

}

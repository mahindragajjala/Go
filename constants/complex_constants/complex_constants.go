package complexconstants

import "fmt"

/*
In Go, complex constants are used to represent complex numbers,
which have a real part and an imaginary part.
*/
func Complex_comstants() {
	//Declaring Untyped Complex Constants
	const x = 3 + 4i // untyped complex constant

	//Assigning Untyped Complex Constants to Variables
	var a complex64 = x  // works: x becomes complex64
	var b complex128 = x // also works: x becomes complex128
	fmt.Println(a, b)

	//Declaring Typed Complex Constants
	const y complex64 = 1 + 2i // typed constant
	fmt.Println(y)
	//Constructing Using complex() Function
	z := complex(2.5, 4.1) // creates 2.5 + 4.1i
	fmt.Println(z)         // Output: (2.5+4.1i)

	//Extracting Parts: real() and imag()
	k := 3 + 4i
	fmt.Println(real(k)) // Output: 3
	fmt.Println(imag(k)) // Output: 4

	const c = 3 + 4i      // untyped complex constant
	var ax complex64 = c  // inferred as complex64
	var bx complex128 = c // inferred as complex128

	fmt.Println("c:", c)
	fmt.Println("a:", ax)
	fmt.Println("b:", bx)
	fmt.Println("Real:", real(c))
	fmt.Println("Imag:", imag(c))
}

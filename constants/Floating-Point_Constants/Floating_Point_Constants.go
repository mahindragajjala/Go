package floatingpointconstants

import "fmt"

/*
In Go, floating-point constants represent decimal (real) numbers
and follow the IEEE-754 standard for floating-point arithmetic.

They can be either typed or untyped, similar to integer constants.

*/
func Floactingpointconstants() {
	const pi = 3.14                 // untyped float constant
	const gravity float64 = 9.80665 // typed float constant
	const sci = 6.022e23            // scientific notation (untyped)
	const avogadro = 6.022e23       // 6.022 × 10^23
	const small = 1.2e-5            // 0.000012

	fmt.Println(pi, gravity, sci, avogadro, small)

}

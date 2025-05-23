package numerictypes

import "fmt"

/*
In engineering, physics, and graphics, complex numbers are used to
represent things like rotation, waveforms, electric circuits, and fractals.

- The real part is like the "horizontal" component.
- The imaginary part is like the "vertical" component.
- Together, they form a point on a 2D plane.
*/
func Complex_numbers() {
	var num complex64 = 2.5 + 3.5i
	fmt.Println("Real part:", real(num))      // Output: 2.5
	fmt.Println("Imaginary part:", imag(num)) // Output: 3.5

}

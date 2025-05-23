package numerictypes

import "fmt"

/*

float32
Size: 32 bits
Precision: ~6–7 decimal digits
Range: ±1.18×10⁻³⁸ to ±3.4×10³⁸

float64
Size: 64 bits
Precision: ~15–16 decimal digits
Range: ±2.23×10⁻³⁰⁸ to ±1.80×10³⁰⁸

*/
func Floating_Point_number() {
	var a float32 = 3.14
	var b float64 = 3.141592653589793

	fmt.Println(a, b)
}

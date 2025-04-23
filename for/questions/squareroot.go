package questions

import (
	"fmt"
	"math"
)

func SquareRoot() {
	// Number whose square root is to be calculated
	num := 16.0

	// Calculate the square root
	sqrt := math.Sqrt(num)

	// Print the result
	fmt.Printf("The square root of %.2f is %.2f\n", num, sqrt)
}

package questions

import "fmt"

/*
Greatest Common Divisor (GCD) Without Modulo

Implement the Euclidean algorithm to find GCD using only subtraction and division.
*/
func gcd(a, b int) int {
	// Loop until the two numbers are equal
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func FindGCD() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	result := gcd(a, b)
	fmt.Printf("The GCD of %d and %d is: %d\n", a, b, result)
}

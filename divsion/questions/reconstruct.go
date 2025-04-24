package questions

import "fmt"

/*

Given divisor, quotient, and remainder, reconstruct the original number.

*/
func reconstructNumber(divisor, quotient, remainder int) int {
	return (divisor * quotient) + remainder
}
func Reconstruct_Number() {
	var divisor, quotient, remainder int

	fmt.Print("Enter divisor: ")
	fmt.Scan(&divisor)

	fmt.Print("Enter quotient: ")
	fmt.Scan(&quotient)

	fmt.Print("Enter remainder: ")
	fmt.Scan(&remainder)

	number := reconstructNumber(divisor, quotient, remainder)
	fmt.Printf("The original number is: %d\n", number)
}

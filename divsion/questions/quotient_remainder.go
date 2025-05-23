package questions

import "fmt"

/*
Find Quotient and Remainder

Write a function that takes two integers a and b, and returns the quotient and remainder of a / b.

*/
func Find_Quotient_and_Remainder(x int, y int) {
	Quotient := x / y
	Remainder := x % y
	fmt.Println(Quotient, Remainder)
}
func Find_Quotient_and_Remainder_input() {
	Find_Quotient_and_Remainder(50, 2)
}

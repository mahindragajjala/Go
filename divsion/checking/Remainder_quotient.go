package checking

import "fmt"

/*
Find Quotient and Remainder

Write a function that takes two integers a and b, and returns the quotient and remainder of a / b.
*/
func Remainder() {
	a := 12            // Dividend
	b := 2             // Divisor
	remainder := a % b //remainder
	fmt.Println(remainder)
}

func Quotient() {
	a := 12
	b := 2
	quotient := a / b //quotient
	fmt.Println(quotient)
}

/*
Dividend=(Divisor×Quotient)+Remainder

Finding All Divisors of a Number:
Any number d such that:

𝑛%𝑑 == 0

is a divisor of n.
*/

package questions

import (
	"fmt"
)

/*
Find Next Multiple
Given a number n and a multiple k, find the smallest number greater than or equal to n that is divisible by k.

You are given:
A number n
A number k (which is the multiple you’re interested in)

You need to find:
The smallest number that is greater than or equal to n and divisible by k.
*/
func findNextMultiple(n, k int) int {
	// Using the math trick to find the smallest number >= n that is divisible by k
	return ((n + k - 1) / k) * k
}

func Smallest_Number() {
	var n, k int

	fmt.Print("Enter number (n): ")
	fmt.Scan(&n)

	fmt.Print("Enter multiple (k): ")
	fmt.Scan(&k)

	result := findNextMultiple(n, k)
	fmt.Printf("The next multiple of %d greater than or equal to %d is: %d\n", k, n, result)
}

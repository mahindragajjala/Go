package questions

import "fmt"

/*
Remainder Pattern

For a given n, print the pattern of remainders when n is divided by numbers from 1 to n.
*/
func Remainder_Pattern() {
	var n int
	fmt.Print("Enter a number (n): ")
	fmt.Scan(&n)

	printRemainderPattern(n)
}
func printRemainderPattern(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", n%i)
	}
	fmt.Println()
}

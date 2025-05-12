package headrecursion

import "fmt"

/*
Reverse a number using head recursion.

    n := 347238947

    lastDigit := n % 10       // Extracts the last digit
    n = n / 10                // Removes the last digit from the number

    fmt.Println("Last digit:", lastDigit)
    fmt.Println("Remaining number:", n)
*/
var ReverseNumber []int

// Head recursion: first recurse, then process
func Reverse_number(n int) {
	if n == 0 {
		return
	}
	Reverse_number(n / 10)                      // Recursive call first
	ReverseNumber = append(ReverseNumber, n%10) // Then process
}

func Reverse_number_input() {
	number := 347238947
	ReverseNumber = []int{} // clear the global slice before use
	Reverse_number(number)

	// Print the reversed number from the digits
	for _, digit := range ReverseNumber {
		fmt.Print(digit)
	}
	fmt.Println()
}

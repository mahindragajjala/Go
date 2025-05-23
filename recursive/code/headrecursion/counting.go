package headrecursion

import "fmt"

/*
Count the number of digits in a number using head recursion.
562346023460
*/

// countDigitsHead recursively counts digits using head recursion
func CountDigitsHead(n int) int {
	if n == 0 {
		return 0
	}
	// Recursive call first (head recursion)
	count := CountDigitsHead(n / 10)

	// Then do the processing
	return count + 1
}

func CountDigitsHead_main() {
	number := 562346023460
	fmt.Printf("Number of digits in %d: %d\n", number, CountDigitsHead(number))
}

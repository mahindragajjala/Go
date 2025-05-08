package code

import "fmt"

// Recursive function to return the nth Fibonacci number
func fibonacci(n int) int {
	fmt.Println("fib:", n)
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func Fibonacci_main() {
	var n int
	fmt.Print("Enter the number of Fibonacci terms: ")
	fmt.Scan(&n)

	fmt.Println("Fibonacci series:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d \n", fibonacci(i))
	}
}

/*
Base Case: If n <= 1, return n directly.

Recursive Case: Otherwise, return the sum of the two previous Fibonacci numbers: fibonacci(n-1) + fibonacci(n-2).
*/

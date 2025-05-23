package questions

import "fmt"

/*
Write a program to print the first 10 Fibonacci numbers using a loop.
➡️ The sequence starts like: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34

➡️ Goal: Understand how values depend on the previous two — loop + logic combo.
*/
/*
0, 1, 1, 2, 3, 5, 8, 13, 21, 34
*/
func Fibonacci_numbers() {
	a, b := 0, 1
	fmt.Println("First 10 Fibonacci numbers:")
	fmt.Print(a, ", ", b)

	for i := 2; i < 10; i++ {
		next := a + b
		fmt.Print(", ", next)
		a = b
		b = next
	}
	fmt.Println()
}

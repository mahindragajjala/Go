package questions

import "fmt"

/*
Write a program to calculate the factorial of a number using a for loop.
Let the number be 5.
(Reminder: 5! = 5 × 4 × 3 × 2 × 1 = 120)
*/
func Factorial() {
	userdata := 50
	Factorial := 1
	for i := userdata; i >= 1; i-- {
		Factorial = Factorial * i
	}
	fmt.Println(Factorial)
}

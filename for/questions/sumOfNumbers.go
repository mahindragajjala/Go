package questions

import "fmt"

/*
Write a program to calculate the sum of numbers from 1 to N using a for loop.
Let N = 50.

➡️ Goal: Learn to accumulate values in a loop (using a variable like sum).
*/
func SumOfNumbers() {
	var Sum int

	for i := 0; i <= 50; i++ {
		Sum += i
	}
	fmt.Println(Sum)
}

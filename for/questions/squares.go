package questions

import "fmt"

/*
Question 10:
Write a program to print the squares of numbers from 1 to 15.
(Example: 1² = 1, 2² = 4, 3² = 9, ..., 15² = 225)

➡️ Goal: Apply mathematical operations inside the loop.
*/
func Squares() {
	for i := 1; i <= 500; i++ {
		fmt.Printf("%v = %v\n", i, i*i)
	}

}

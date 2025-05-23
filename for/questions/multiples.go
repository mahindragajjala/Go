package questions

import "fmt"

/*
Write a program to print all multiples of 3 or 5 below 100 using a for loop.

➡️ Expected Output: 3, 5, 6, 9, 10, 12, 15, ...

➡️ Goal: Work with multiple conditions to filter numbers in a range.
*/
func PrintAllMultiple() {
	var multiples []int
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			multiples = append(multiples, i)
		}
	}
	fmt.Println(multiples)
}

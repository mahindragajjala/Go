package questions

import "fmt"

/*
Write a program to print the following pattern using nested loops for N = 5:

5
5 4
5 4 3
5 4 3 2
5 4 3 2 1

➡️ Goal: Combine reverse number logic with nested loops for pattern printing.
*/
func Reversed_numbered_pattern() {
	var slice []int
	for i := 5; i >= 1; i-- {
		slice = append(slice, i)
		for _, value := range slice {
			fmt.Print(value)
		}
		fmt.Println()
	}
}

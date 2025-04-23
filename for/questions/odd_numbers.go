package questions

import "fmt"

/*
Write a program to print all odd numbers from 1 to 30 using a for loop.

➡️ Goal: Strengthen your understanding of conditions (if) and loop filtering.
*/
func ODD_numbers() {
	ODD_numbers := []int{}
	for i := 1; i <= 30; i++ {

		if i%2 != 0 {
			ODD_numbers = append(ODD_numbers, i)
		}
	}
	fmt.Println(ODD_numbers)
}

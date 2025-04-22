package questions

import "fmt"

/*
Write a program to print even numbers from 1 to 20 using a for loop.
➡️ Goal: Learn how to use conditions (if) inside a for loop to filter values.
*/
var Even []int

func Even_numbers() {
	for i := 1; i <= 20; i++ {
		//check remainder..:)
		if i%2 == 0 {
			Even = append(Even, i)
		}
	}
	fmt.Println(Even)
}

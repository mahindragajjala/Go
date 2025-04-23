package questions

import "fmt"

/*
Check if Number is Divisible

Write a function to check if a number n is divisible by k.
*/

func Check_Divisible(n int, k int) {
	if n%k == 0 {
		fmt.Println("divisible")
	} else {
		fmt.Println("not divisible")
	}
}
func Check_Divisible_input() {
	Check_Divisible(25, 3)
}

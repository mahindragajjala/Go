package questions

import "fmt"

/*
Write a program to count the number of digits in a given number using a loop.
Let the number be 987654.

➡️ Expected Output: 6 (because it has 6 digits)

➡️ Goal: Strengthen loop control and digit-based logic using / or string length approach.
*/
func Count_numbers(MN int) int {
	var MOBILE_NUMBER int = MN
	fmt.Println("Your mobile was : ", MOBILE_NUMBER)
	var count int

	for MN > 0 {
		MN = MN / 10
		count++
	}
	return count
}
func Mobile_number() {
	count := Count_numbers(7396522068)
	fmt.Printf("Your Mobile Number Has : %v", count)
}

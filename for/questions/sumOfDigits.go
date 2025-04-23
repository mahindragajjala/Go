package questions

import "fmt"

/*
Write a program to print the sum of the digits of a given number using a for loop (or a loop logic).
Let the number be 567.

➡️ Expected Output: 5 + 6 + 7 = 18

➡️ Goal: Practice loop logic with digit extraction using % and /.
*/
func SumOfDigits() {
	number := 567
	//original := number
	sum := 0

	// For storing the digits for display like 5 + 6 + 7
	digits := []int{}

	for number > 0 {
		digit := number % 10
		sum += digit
		digits = append([]int{digit}, digits...) // Prepend to preserve order
		number /= 10
	}

	// Print digits and sum
	for i, digit := range digits {
		fmt.Print(digit)
		if i != len(digits)-1 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

package questions

import "fmt"

/*
Write a program to find the smallest number in an array using a for loop.
Let the array be: [18, 4, 25, 9, 2, 31].

➡️ Expected Output: 2 (the smallest number)

➡️ Goal: Similar to the max logic, now tracking the minimum value during iteration.
*/
func FindSmallestNo() {
	array := [6]int{18, 4, 25, 9, 2, 31}
	smallest := array[0] // Initialize with first element

	for _, value := range array[1:] {
		if value < smallest {
			smallest = value
		}
	}

	fmt.Println(smallest) // Expected Output: 2
}

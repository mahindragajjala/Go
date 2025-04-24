package questions

import "fmt"

/*
	Write a program to find the largest number in an array of integers using a for loop.
	Let the array be: [5, 12, 3, 20, 8, 7].

	➡️ Expected Output: 20 (the largest number)
	➡️ Goal: Practice iterating through an array and updating a maximum value.
*/
func Largest_Number() {
	array := [6]int{5, 12, 3, 20, 8, 7}
	//highestno := array[0]
	var variable int

	for i := 1; i < len(array); i++ {
		if array[i] > variable {
			variable = array[i]
		}
	}

	//fmt.Println("Largest number is:", highestno)
	fmt.Println("Largest number is:", variable)
}

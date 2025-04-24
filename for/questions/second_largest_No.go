package questions

import "fmt"

/*
Write a program to find the second largest number in an array using a for loop.
Let the array be: [12, 35, 1, 10, 34, 1]

➡️ Expected Output: 34

➡️ Goal: Practice comparing and tracking multiple values using loop logic.
*/
func Second_largest_number() {
	array := [6]int{12, 35, 1, 10, 34, 1}
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
	fmt.Println("second longest number :", array[len(array)-2])
}

package questions

import "fmt"

/*
Write a program to find the second largest number in an array using a for loop.
Let the array be: [12, 35, 1, 10, 34, 1]

➡️ Expected Output: 34

➡️ Goal: Practice comparing and tracking multiple values using loop logic.
*/
func Second_largest_number() {
	array := []int{12, 35, 1, 10, 34, 1}
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	fmt.Println(array)

	//finding second largest
	fmt.Println(array[n-2])
}

package questions

import "fmt"

/*
Write a program to move all zeroes to the end of an array using a for loop.
Let the array be: [0, 1, 0, 3, 12]

➡️ Expected Output: [1, 3, 12, 0, 0]

➡️ Goal: Strengthen array manipulation and element shifting logic.
*/
func Move_Zeros_to_End() {
	array := []int{0, 1, 0, 3, 12}
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if array[j] == 0 {
				array[j], array[j+1] = array[j+1], 0
			}
		}
	}
	fmt.Println(array)
}

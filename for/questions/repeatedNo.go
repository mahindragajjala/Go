package questions

import "fmt"

/*
Write a program to count how many times a given number appears in an array.
Let the array be: [1, 2, 3, 4, 2, 5, 2, 6]
Let the number to search be: 2

➡️ Expected Output: 2 appears 3 times

➡️ Goal: Strengthen your loop logic and conditional counting.
*/
func Count_Numbers(x int) {
	array := [8]int{1, 2, 3, 4, 2, 5, 2, 6}
	var count int
	for i := 0; i < len(array); i++ {
		if x == array[i] {
			count++
		}
	}
	fmt.Println(count)
}
func Count_numbers_input() {
	Count_Numbers(2)
}

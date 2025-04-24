package questions

import "fmt"

/*
Write a program to count how many positive, negative, and zero values are in an array using a for loop.
Let the array be: [4, -3, 0, 7, -1, 0, 9]

Expected Output:
Positive: 3
Negative: 2
Zero: 2

Goal: Use conditionals inside loops to classify data.
*/
func Count_Positive_negative() {
	array := []int{4, -3, 0, 7, -1, 0, 9}
	var Positive int
	var Negative int
	var Zeros int
	for i := 0; i < len(array); i++ {
		if array[i] < 0 {
			Negative++
		} else if array[i] > 0 {
			Positive++
		} else {
			Zeros++
		}
	}
	fmt.Println(Positive)
	fmt.Println(Negative)
	fmt.Println(Zeros)
}

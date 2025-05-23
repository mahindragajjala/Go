package questions

import "fmt"

/*
Write a program to merge two arrays into a third array using a for loop.Let the arrays be:A = [1, 3, 5]B = [2, 4, 6]➡️ Expected Output: [1, 3, 5, 2, 4, 6]
➡️ Goal: Practice looping over two data sources and storing results in a new one.
*/
func Merge_two_arrays_into_third_array() {
	a := [3]int{1, 3, 5}
	b := [6]int{2, 4, 6}
	var c []int
	for _, value := range a {
		c = append(c, value)
	}
	for _, value := range b {
		c = append(c, value)
	}
	fmt.Println(c)
}

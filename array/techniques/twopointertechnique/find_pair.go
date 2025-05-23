package twopointertechnique

import "fmt"

/*
1. Find Pair with Target Sum in Sorted Array
   Given a sorted array, return true if there exists a pair that sums up to a target value.

*/
func Find_Pair(array [5]int, target int) bool {
	left := 0
	right := len(array) - 1

	for left < right {
		if array[left]+array[right] == target {
			fmt.Println(array[left], array[right])
			left++
			right--
			return true
		}
	}
	return false
}
func Find_Pair_Data(array [5]int) {

	left := 0
	right := len(array) - 1

	for left < right {
		fmt.Println(array[left], array[right])
		left++
		right--

	}
}

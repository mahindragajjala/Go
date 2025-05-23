package questions

import "fmt"

/*
Write a program to find all pairs in an array whose sum is equal to a given number using a for loop.
Let the array be: [1, 5, 7, -1, 5] and the target sum be 6.

➡️ Expected Output:

scss
Copy
Edit
(1, 5)
(7, -1)
(1, 5)  // (depends on handling duplicate 5s)
➡️ Goal: Practice nested loops and pair combination logic.
*/
func SumofPairs() {
	array := []int{1, 5, 7, -1, 5}
	n := len(array)
	TargetSum := 6
	var NestedSlice [][]int

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if array[i]+array[j] == TargetSum {
				NestedSlice = append(NestedSlice, []int{array[i], array[j]})
			}
		}
	}

	for _, pair := range NestedSlice {
		fmt.Println(pair)
	}
}

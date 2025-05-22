package traversal

import "fmt"

/*
	Finding Maximum / Minimum Element
	Track and update while traversing.

	🔹 Minimum (min) Value
	The smallest number in a list or range.
	Example:
	In the set [4, 2, 9, 1, 7], the min value is 1.

	🔹 Maximum (max) Value
	The largest number in a list or range.
	Example:
	In the set [4, 2, 9, 1, 7], the max value is 9.
*/
func FindMaxMini() {
	Max := 0
	Mini := 0
	arr := []int{3, 1, 6, 3, 7, 9}
	for i := 0; i < len(arr); i++ {
		if Max == 0 && Mini == 0 {
			Max = arr[i]
			Mini = arr[i]
		} else {
			if arr[i] < Max {
				Max = arr[i]
			} else {
				Mini = arr[i]
			}
		}
	}
	fmt.Println(Max, Mini)
}

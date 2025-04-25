package questions

import "fmt"

/*
ascending - 5, 2, 9, 1, 3
1 - 2,5,9,1,3
    2,5,9,1,3
	2,5,1,9,3
	2,5,1,3,9
2 - 2,5,1,3,9
	2,1,5,3,9
	2,1,3,5,9
	2,1,3,5,9
3 - 1,2,3,5,9

*/

func Ascending_order() {
	arr := []int{5, 2, 9, 1, 3}

	// Bubble sort (ascending)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted in ascending order:", arr)
}

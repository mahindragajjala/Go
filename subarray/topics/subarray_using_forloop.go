package topics

import "fmt"

func Subarray() {
	// Sample input array
	arr := []int{1, 2, 3}
	n := len(arr)

	fmt.Println("All Subarrays:")
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// Print subarray arr[i..j]
			fmt.Print("[")
			for k := i; k <= j; k++ {
				fmt.Print(arr[k])
				if k != j {
					fmt.Print(", ")
				}
			}
			fmt.Println("]")
		}
	}
}

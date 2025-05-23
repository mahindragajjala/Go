package questions

import "fmt"

func Counting_swaps_in_sorting() {
	array := []int{6, 8, 4, 2, 1, 27, 9, 6}
	n := len(array)
	var count int
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				count++
			}
		}
	}
	fmt.Println(array)
	fmt.Println(count)
}

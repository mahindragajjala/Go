package questions

import "fmt"

func Add_flag_sorting() {
	array := []int{7, 8, 1, 2, 6, 2, 3}
	n := len(array)

	for i := 0; i < n-1; i++ {
		flag := 0 // reset flag at the beginning of each outer loop iteration
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				// swap elements
				array[j], array[j+1] = array[j+1], array[j]
				flag = 1 // mark that a swap occurred
			}
		}
		if flag == 0 {
			break // if no swaps, array is sorted
		}
	}

	fmt.Println(array)
}

package questions

import "fmt"

/*
Sort Slice of Float64
Modify your Bubble Sort to sort a slice of float64 numbers.
*/
func Float_number_slice() {
	array := []float64{2.24, 4.2, 1, 2, 8.2, 5.2, 3.6}
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}

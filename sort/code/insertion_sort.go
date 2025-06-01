package code

import "fmt"

/*
For an array of size n:

Start from index i = 1 (because a single element at index 0 is already "sorted").
Store the current value key = arr[i].
Compare key with elements in the sorted part (arr[0...i-1]).
Shift elements rightward if they are greater than key.
Insert key in its correct position.

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]         // key = arr[i]
		j := i - 1            // j = i - 1

		for j >= 0 && arr[j] > key { // while j >= 0 and arr[j] > key
			arr[j+1] = arr[j]        // arr[j+1] = arr[j]
			j = j - 1                // j = j - 1
		}
		arr[j+1] = key         // arr[j+1] = key
	}
}

*/

func Insertion_sort() {
	array := []int{85, 12, 59, 45, 72, 51}
	n := len(array)
	for i := 1; i < n-1; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
	}
}
func Insertion_sort_another_method() {
	array := []int{5, 4, 10, 1, 6, 2}
	n := len(array)
	for i := 1; i < n; i++ {
		temp := array[i]
		j := i - 1
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = temp
	}
	fmt.Println(array)
}

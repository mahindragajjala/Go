package code

import "fmt"

/*


algorithm:
For an array of size n:
Start from index i = 0.
Find the index of the minimum element in the subarray from i to n-1.
Swap the element at index i with the minimum element.
Repeat step 1 to 3 for i = 1 to n-2.

for i = 0 to n-1
    min = i
    for j = i+1 to n
        if arr[j] < arr[min]
            min = j
    swap arr[i] and arr[min]

*/
func Selection_sort() {
	array := []int{20, 12, 10, 15, 2}
	selectionSort(array)
	fmt.Println(array)
}
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

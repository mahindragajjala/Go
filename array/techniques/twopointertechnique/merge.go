package twopointertechnique

/*
5. Merge Two Sorted Arrays into One
   *Given two sorted arrays, merge them into one sorted array using TPT.*

*/
func Merge_sorted_arrays(array [5]int, array2 [5]int) [10]int {
	newarray := [10]int{}

	// Copy elements from array
	for i := 0; i < len(array); i++ {
		newarray[i] = array[i]
	}

	// Copy elements from array2 starting at index 5
	for j := 0; j < len(array2); j++ {
		newarray[j+5] = array2[j]
	}

	return newarray
}

func MergeSortedArrays(arr1, arr2 []int) []int {
	i, j := 0, 0
	merged := make([]int, 0, len(arr1)+len(arr2))

	// Traverse both arrays using two pointers
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	// Append remaining elements
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}

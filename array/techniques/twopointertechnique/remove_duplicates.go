package twopointertechnique

import "fmt"

/*
Function RemoveDuplicates(array):
    If array is empty:
        Return 0

    Initialize slow = 0

    For fast = 1 to length of array - 1:
        If array[fast] ≠ array[slow]:
            Increment slow by 1
            array[slow] = array[fast]

    Return slow + 1

*/
// RemoveDuplicates removes duplicates in-place from a sorted slice and returns the new length.
func RemoveDuplicates(array []int) int {
	if len(array) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(array); j++ {
		if array[j] != array[i] {
			i++
			array[i] = array[j]
		}
	}
	return i + 1 // new length
}

func RD() {
	arr := []int{2, 2, 3, 3, 5, 5, 7, 8}
	newLen := RemoveDuplicates(arr)
	fmt.Println("After removing duplicates:", arr[:newLen])
}

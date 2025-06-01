package tailrecursion

func IsSorted(arr []int, index int) bool {
	if index == len(arr)-1 {
		return true
	}
	if arr[index] > arr[index+1] {
		return false
	}
	return IsSorted(arr, index+1)
}

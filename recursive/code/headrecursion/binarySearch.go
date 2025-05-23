package headrecursion

func BinarySearch(arr []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == target {
		return mid
	} else if target < arr[mid] {
		return BinarySearch(arr, left, mid-1, target)
	} else {
		return BinarySearch(arr, mid+1, right, target)
	}
}

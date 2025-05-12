package headrecursion

func MaxInArray(arr []int, index int) int {
	if index == len(arr)-1 {
		return arr[index]
	}
	max := MaxInArray(arr, index+1)
	if arr[index] > max {
		return arr[index]
	}
	return max
}

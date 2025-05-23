package tailrecursion

func RemoveDuplicatesTail(arr []int, index int, seen map[int]bool, result []int) []int {
	if index == len(arr) {
		return result
	}
	if !seen[arr[index]] {
		seen[arr[index]] = true
		result = append(result, arr[index])
	}
	return RemoveDuplicatesTail(arr, index+1, seen, result)
}

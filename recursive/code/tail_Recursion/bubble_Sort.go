package tailrecursion

func BubbleSortTail(arr []int, n int) {
	if n == 1 {
		return
	}
	bubblePass(arr, 0, n)
	BubbleSortTail(arr, n-1)
}

func bubblePass(arr []int, i, n int) {
	if i == n-1 {
		return
	}
	if arr[i] > arr[i+1] {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	bubblePass(arr, i+1, n)
}

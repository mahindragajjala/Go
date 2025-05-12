package tailrecursion

import "fmt"

func PrintArrayReverse(arr []int, index int) {
	if index < 0 {
		return
	}
	fmt.Println(arr[index])
	PrintArrayReverse(arr, index-1)
}

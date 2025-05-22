package topics

import "fmt"

func SubarrayWithForloop() {
	arr := []int{5, 7, 8, 2, 1, 8, 4}

	// Two loops: one for start index, one for end index
	for start := 0; start < len(arr); start++ {
		for end := start + 1; end <= len(arr); end++ {
			subarray := arr[start:end]
			fmt.Println(subarray)
		}
	}
}

package traversal

import "fmt"

func Comparing(array [5]int, value int) {
	for i := 1; i <= len(array); i++ {
		if array[i] == value {
			fmt.Println("exist:", value)
		}
	}
}

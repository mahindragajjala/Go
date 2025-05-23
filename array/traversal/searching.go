package traversal

import "fmt"

func Searching(array [5]int, value int) {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			fmt.Println("Exist:", value)
		}
	}
}

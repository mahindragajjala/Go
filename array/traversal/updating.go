package traversal

import "fmt"

func Updating(array [5]int, index int, replacevalue int) {
	for i := 0; i < len(array); i++ {
		if i == index {
			array[i] = replacevalue
		}
	}
	fmt.Println("NewArray:", array)
}

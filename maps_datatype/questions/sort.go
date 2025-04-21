package questions

import (
	"fmt"
)

/*
Write a Go program to sort a map by keys and print the sorted output.
*/
func SortingMap() {
	originalMap_data := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	for _, value := range originalMap_data {
		fmt.Println(value)
	}
}

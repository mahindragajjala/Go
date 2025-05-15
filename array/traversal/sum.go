package traversal

import "fmt"

func Sum(array [5]int) {
	var Sum int
	for i := 0; i < len(array); i++ {
		Sum = Sum + array[i]
	}
	fmt.Println(Sum)
}

package traversal

import "fmt"

func Reverse_traversal(array [5]int) {
	for i := len(array) - 1; i >= 0; i-- { // 4 >= 0 = true
		fmt.Println("Reverse array : ", array[i])
	}
	fmt.Println(array[0])
}

package questions

import "fmt"

func Reverse_Array() {
	array := [5]int{1, 2, 3, 4, 5}
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5, 6}

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j+1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println("slice:", slice)
}

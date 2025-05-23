package questions

import (
	"fmt"
)

func Decending_order() {
	array := []int{12, 35, 1, 10, 34, 1}
	n := len(array) //5
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if array[j] < array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	fmt.Println("Decending order", array)

}

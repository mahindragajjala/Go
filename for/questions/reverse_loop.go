package questions

import "fmt"

func Reverse_Printing() {
	var slice []int
	for i := 5; i >= 1; i-- {
		slice = append(slice, i)
		fmt.Println(slice)
	}

}

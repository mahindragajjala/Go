package problems

import "fmt"

//Write code that prints the sum of each row in the slice.

func SumOfrows() {
	slicedata := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var sum int
	for _, value := range slicedata[1] {
		sum += value
		fmt.Println(value)
	}
	fmt.Println("sum-->", sum)
}

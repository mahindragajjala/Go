package pipeline

import "fmt"

func SQUARE_PIPELINE() {
	data := []int{1, 2}
	num := inserinput(data)
	squarevalue := squares(num)
	Print(squarevalue)
}
func inserinput(nums []int) <-chan int {
	data := make(chan int)
	for v := range nums {
		data <- v
	}
	return data
}
func squares(data <-chan int) chan int {
	ch := make(chan int)
	for v := range data {
		square := v * v
		ch <- square
	}
	return ch
}
func Print(sq chan int) {
	for v := range sq {
		fmt.Println(v)
	}
}

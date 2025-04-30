package code

import "fmt"

var slice []int

func Fibonacci_series(a int, b int) {
	if len(slice) >= 5 {
		return
	}
	slice = append(slice, a)
	Fibonacci_series(b, a+b)
}

func Fibonacci() {
	slice = []int{}
	Fibonacci_series(0, 1)
	fmt.Println(slice)
}

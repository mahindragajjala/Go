package code

import "fmt"

var slicedata []int

func fib(a int, b int) int { //func fib(0,1)///func fib(1,1)
	if len(slicedata) == 5 { //false///false
		return 0
	}
	if len(slicedata) == 0 { //true slicedata = []int{0,1} ///false
		slicedata = append(slicedata, a, b)
	}
	addition := a + b                       // 0 + 1 = 1 /// 1 + 1 = 2
	slicedata = append(slicedata, addition) //slicedata = []int{0,1,1} /// slicedata = []int{0,1,1,2}
	fmt.Println(slicedata)
	return fib(b, addition) //func fib(1,1) /// func fib(1,2)
}
func Fib_main() {
	fib(0, 1)
}

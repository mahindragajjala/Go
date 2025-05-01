package code

import "fmt"

func Tail_recusion() {
	/*
		Tail recursion is a special kind of recursion where the recursive call is the "last thing executed" in the function.

		In such cases, there's no need to keep the current function's "state on the call stack", which allows compilers or interpreters to optimize and reuse stack frames — this is called Tail Call Optimization (TCO).
	*/
	var factorial int = 1
	factorialdata := Test_factorial(5, factorial)
	fmt.Println(factorialdata)
}
func Test_factorial(x int, factorial int) int {
	if x == 1 {
		return factorial
	}
	factorial *= x
	return Test_factorial(x-1, factorial)
}

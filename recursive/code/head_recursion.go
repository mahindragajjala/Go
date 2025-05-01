package code

import "fmt"

/*
Head recursion is a type of recursion where the recursive call happens before any other operations in the function.

🔁 The function calls itself first, and then does computation after returning from the recursive call
5!
5 * 4 * 3 * 2 * 1

n * factorial(5-1 = 4) * factorial(5-1 = 4) * factorial(4 - 1 = 3) *factorial(3-1 = 2) * factorial(2-1 = 1) *  factorial(if n == 1 = exist && return n)

Function HeadRec(n):
    If n == 0:
        Return
    Call HeadRec(n - 1)      // Recursive call first (head recursion)
    Print n                  // Work happens after the recursive call
*/
func Head_recursive(x int) int {
	if x == 1 {
		return x
	}
	return x * Head_recursive(x-1)
}
func Test_Head_recursive() {
	factorial := Head_recursive(5)
	fmt.Println(factorial)
}

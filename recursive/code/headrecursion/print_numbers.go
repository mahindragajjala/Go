package headrecursion

import "fmt"

/*
Print numbers from 1 to n using head recursion
n := 6
1,2,3,4,5,6
base condition = 1
*/
func Printing_numbers_using_recursion(n int) int {
	if n == 1 {
		fmt.Println(n)
		return 1
	} else {
		Printing_numbers_using_recursion(n - 1)
		fmt.Println(n)
		return n
	}
}
func Printing_numbers_using_recursion_input() {
	Printing_numbers_using_recursion(6)
}

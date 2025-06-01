package headrecursion

import "fmt"

/*
	Find the sum of digits of a number using head recursion.
	6 + 5 + 4 + 3 + 2 +  = 21
	n := 6
	6 -> sum + n-1 + sum + n-1 + .... 1 == if n == 1{
	retrun 1
	}
	base condition 1
*/
var sum int

func SumOfDigits(n int) int {
	if n == 1 {
		sum += n
		return sum
	}
	SumOfDigits(n - 1)
	sum += n
	return sum
}
func SumOfDigits_input() {
	sum := SumOfDigits(6)
	fmt.Println("sum : ", sum)
}

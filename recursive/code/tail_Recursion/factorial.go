package tailrecursion

func FactorialTail(n, acc int) int {
	if n == 0 {
		return acc
	}
	return FactorialTail(n-1, acc*n)
}

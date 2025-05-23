package tailrecursion

func FibonacciTail(n, a, b int) int {
	if n == 0 {
		return a
	}
	return FibonacciTail(n-1, b, a+b)
}

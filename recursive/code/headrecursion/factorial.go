package headrecursion

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := Factorial(n - 1)
	return n * result
}

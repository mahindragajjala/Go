package recursivefunction

/*
Functions that call themselves.
*/
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// see the recursive folder for more information:)

package tailrecursion

func PowerTail(x, n, acc int) int {
	if n == 0 {
		return acc
	}
	return PowerTail(x, n-1, acc*x)
}

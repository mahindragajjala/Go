package tailrecursion

func SumDigits(n, acc int) int {
	if n == 0 {
		return acc
	}
	return SumDigits(n/10, acc+(n%10))
}

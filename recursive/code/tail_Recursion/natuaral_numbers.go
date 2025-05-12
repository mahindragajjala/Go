package tailrecursion

func SumNaturalTail(n, acc int) int {
	if n == 0 {
		return acc
	}
	return SumNaturalTail(n-1, acc+n)
}

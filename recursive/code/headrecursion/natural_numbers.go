package headrecursion

func SumNatural(n int) int {
	if n == 0 {
		return 0
	}
	sum := SumNatural(n - 1)
	return sum + n
}

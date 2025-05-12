package tailrecursion

func MultiplyTail(a, b, acc int) int {
	if b == 0 {
		return acc
	}
	return MultiplyTail(a, b-1, acc+a)
}

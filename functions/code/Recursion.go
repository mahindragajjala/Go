package code

func Recursion(a int) int {
	if a == 0 {
		return 1
	}
	return Recursion(a - 1)
}

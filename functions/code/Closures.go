package code

func Closures() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

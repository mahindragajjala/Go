package code

func Function_as_Return_Value(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

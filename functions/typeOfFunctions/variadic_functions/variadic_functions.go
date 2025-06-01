package variadicfunctions

/*
A variadic function is a function that can accept a variable number of arguments.

In many languages (including Go, Python, JavaScript, etc.), this feature allows you to pass zero or more values of a specified type.
*/
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

/*
How it Works:
The ...int means zero or more integers.

Inside the function, nums is treated like a slice ([]int).
*/

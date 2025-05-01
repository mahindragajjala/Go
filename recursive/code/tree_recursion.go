package code

/*
What Is Tree Recursion?
In linear recursion, a function calls itself once.
In tree recursion, a function calls itself multiple times, leading to a tree of recursive calls.
Use it when:

The problem naturally branches (e.g., binary trees, permutations, combinations)
Each call needs to explore multiple possibilities.
*/
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

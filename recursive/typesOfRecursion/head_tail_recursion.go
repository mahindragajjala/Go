package typesOfRecursion

//Tail vs Head Recursion (Optional, for deeper understanding)
/*
Definition: The recursive call is the last operation in the function.
Key Point: No computation is pending after the recursive call.
Optimization: Can be optimized by the compiler using Tail Call Optimization (TCO), meaning it doesn't add a new stack frame.

*/
func TailSum(n, acc int) int {
	if n == 0 {
		return acc
	}
	return TailSum(n-1, acc+n)
}

/*
tailFactorial(5, 1)
tailFactorial(4, 5)
tailFactorial(3, 20)
tailFactorial(2, 60)
tailFactorial(1, 120)
return 120
*/
/*
Definition: The recursive call is the first operation in the function.
Key Point: Some computation is pending after the recursive call returns.
Optimization: Cannot be optimized as easily — needs stack memory.
*/
func HeadSum(n int) int {
	if n == 0 {
		return 0
	}
	return n + HeadSum(n-1)
}

/*
headFactorial(5)
→ headFactorial(4)
→ headFactorial(3)
→ headFactorial(2)
→ headFactorial(1)
→ return 1
→ return 2
→ return 6
→ return 24
→ return 120
*/

/*
Tail recursion: Work is done before the next call. Nothing left after the call.
Head recursion: Work is done after the next call. You wait for the recursive result.
*/

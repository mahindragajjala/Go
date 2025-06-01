package typesOfRecursion

import "fmt"

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
func Tree_recursion() {
	result := fibonacci(5)
	fmt.Println(result)
}

/*
// C++ program to show Tree Recursion
#include <iostream>
using namespace std;

// Recursive function
void fun(int n)
{
	if (n > 0)
	{
		cout << " " << n;

		// Calling once
		fun(n - 1);

		// Calling twice
		fun(n - 1);
	}
}

// Driver code
int main()
{
	fun(3);
	return 0;
}

// This code is contributed by shivanisinghss2110

*/

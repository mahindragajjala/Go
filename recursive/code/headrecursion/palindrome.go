package headrecursion

import "fmt"

func IsPalindromeHelper(n, temp int) (int, bool) {
	if n == 0 {
		return temp, true
	}
	reversed, _ := IsPalindromeHelper(n/10, temp)
	reversed = reversed*10 + n%10
	return reversed, reversed == temp
}

func IsPalindrome(n int) bool {
	reversed, isPal := IsPalindromeHelper(n, n)
	fmt.Println(reversed)
	return isPal
}

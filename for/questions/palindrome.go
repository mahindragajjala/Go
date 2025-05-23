package questions

import "fmt"

/*
Write a program to check if a number is a palindrome using a loop.
Let the number be 121.
➡️ A number is a palindrome if it reads the same backward (e.g., 121, 1221, 1331).

➡️ Goal: Combine reverse logic with a comparison for real-world validation problems.
*/

func Palindrome(x int) {
	var slice []int
	actual := x
	var reversed int
	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x = x / 10
	}

	fmt.Println(slice)

	if actual == reversed {
		fmt.Println("palindrome...")
	} else {
		fmt.Println("not palindrome...")
	}
}
func Palindrome_userinput() {
	Palindrome(1221)
}

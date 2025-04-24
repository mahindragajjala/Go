package questions

import "fmt"

/*
Write a program to print the following number pattern for N = 5:

1
1 2
1 2 3
1 2 3 4
1 2 3 4 5
*/
func Numbered_patterns() {
	n := 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

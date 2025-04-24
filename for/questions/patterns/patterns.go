package questions

import "fmt"

/*
Write a program to print a pattern of stars (*) in the shape of a right-angled triangle.
For N = 5, the output should be:
*
**
***
****
*****
*/
func PrintPattern() {
	N := 5

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

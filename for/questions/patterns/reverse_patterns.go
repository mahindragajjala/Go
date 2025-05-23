package questions

import "fmt"

/*
Write a program to print an inverted right-angled triangle pattern of stars.
For N = 5, the output should be:

*****
****
***
**
*

*/
func Reverse_patterns() {
	n := 5
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

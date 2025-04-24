package questions

import "fmt"

/*
Write a program to print this pattern for N = 5:

mathematical

E
E D
E D C
E D C B
E D C B A
➡️ Goal: Combine reverse alphabet printing with nested loops.
*/
func Reversed_letter_pattern() {
	for i := 1; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", 'E'-j)
		}
		fmt.Println()
	}
}

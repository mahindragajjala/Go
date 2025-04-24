package questions

import "fmt"

/*
Write a program to print this pattern for N = 5:

A
A B
A B C
A B C D
A B C D E
➡️ Goal: Use character manipulation (char) in loops to create alphabet patterns.
*/
func Letter_pattern() {
	N := 5

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", 'A'+j)
		}
		fmt.Println()
	}
}

/*
🔁 Outer Loop: i = 0
Enters inner loop: j = 0

'A' + 0 → 'A' → prints A

A

🔁 Outer Loop: i = 1
Enters inner loop:

j = 0 → 'A' + 0 → A

j = 1 → 'A' + 1 → B

A
A B

🔁 Outer Loop: i = 2
Inner loop:

j = 0 → 'A' + 0 → A

j = 1 → 'A' + 1 → B

j = 2 → 'A' + 2 → C

A
A B
A B C

🔁 Outer Loop: i = 3
Inner loop:

j = 0 → A

j = 1 → B

j = 2 → C

j = 3 → D

A
A B
A B C
A B C D

🔁 Outer Loop: i = 4
Inner loop:

j = 0 → A

j = 1 → B

j = 2 → C

j = 3 → D

j = 4 → E

A
A B
A B C
A B C D
A B C D E

*/

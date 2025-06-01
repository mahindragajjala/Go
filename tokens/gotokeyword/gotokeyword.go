package gotokeyword

import "fmt"

/*
the goto keyword is used to jump to a labeled statement within the same function.

It's like saying: “Hey, skip all the lines in between and go directly to that label.”
Error handling/cleanup	    ✅ Yes
Exiting nested loops	    ✅ Yes
Simple conditionals/loops	❌ No
Regular flow control	    ❌ No
*/
func GotoKeyword() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 6 {
			goto TheEnd
		}
	}
TheEnd:
	fmt.Println("the end...")
}

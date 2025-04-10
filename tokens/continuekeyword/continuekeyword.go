package continuekeyword

import "fmt"

/*
continue - it's used inside loops and is pretty straightforward once you get the hang of it.
*/
func Continue_keyword() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}

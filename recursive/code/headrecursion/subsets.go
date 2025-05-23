package headrecursion

import "fmt"

func SubsetsHelper(s, current string, index int) {
	if index == len(s) {
		fmt.Println(current)
		return
	}
	SubsetsHelper(s, current, index+1)                  // exclude current character
	SubsetsHelper(s, current+string(s[index]), index+1) // include current character
}

func Subsets(s string) {
	SubsetsHelper(s, "", 0)
}

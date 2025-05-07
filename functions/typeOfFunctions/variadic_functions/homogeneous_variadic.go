package variadicfunctions

import "fmt"

/*
All arguments are of the same type
*/
func PrintAll(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

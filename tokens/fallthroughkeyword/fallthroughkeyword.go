package fallthroughkeyword

import "fmt"

/*
In Go, fallthrough is a keyword used in a switch statement to force the execution to
continue to the next case, regardless of whether that case matches the condition.

fallthrough is a control flow keyword that, when used as the last statement of a case
block in a switch, causes the program to continue executing the next case block
unconditionally, without checking its condition.

*/
func Fallthroughfunc() {
	number := 2

	switch number {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("Other number")
	}
}

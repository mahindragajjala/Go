package booleanconstants

import "fmt"

/*
boolean constants are values that represent true or false.
These are the only two possible values for a variable of type bool
*/
func Bool_constants() {
	var a bool = true
	var b = false
	c := true

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}

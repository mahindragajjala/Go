package inputpassing

import "fmt"

func PassByValue() {
	fmt.Println("passbyvalue-->")
	var a int = 25
	Taking_passbyvalue(a)
}
func Taking_passbyvalue(a int) {
	fmt.Println(a)
}

/*
Copy semantics
The integer a in PassByValue is copied into a new local variable a in Taking_passbyvalue.

Isolation
Mutating a inside the function does not affect the caller’s a.

When to use

You don’t need to modify the caller’s variable.

You want to keep the called function from accidentally changing your data.

The value is small/simple (e.g., int, float, small struct).
*/

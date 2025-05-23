package inputpassing

import "fmt"

func PassByReference() {
	fmt.Println("passbyreference-->")
	var a int = 100
	Taking_passbyreference(&a)
}
func Taking_passbyreference(a *int) {
	fmt.Println(*a)

}

/*
Pointer semantics
You pass the address of the variable (&a) so inside the function you get a pointer *int.

Shared access
Dereferencing (*a) reads or writes the original variable’s memory.

When to use

You need the function to modify the caller’s variable.

You want to avoid copying a large data structure (e.g., big slice header, large struct).

You need to share state between caller and callee.
*/

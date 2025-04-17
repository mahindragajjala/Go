package code

import "fmt"

/* Interface Values and Nil
Understand nil interface traps.
Experiment with nil interfaces and their behavior.


In Go, an interface is made of two components:

Type — the actual concrete type stored.
Value — the concrete value.

When both are nil, the interface is truly nil.
But if the Type is set but Value is nil, the interface is not nil.
That’s where the trap lies.
*/

//Both Type and Value are nil
func Both_type_and_value_areNil() {
	var i interface{} = nil // Both type and value are nil

	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is NOT nil")
	}
}

type Person struct {
	Name string
}

func Type_is_set_but_Value_is_nil() {
	var p *Person = nil   // p is a nil pointer
	var i interface{} = p // Assigning nil pointer to interface

	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is NOT nil")
	}

	var j interface{} = "aaaa" // Assigning nil pointer to interface

	if j == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is NOT nil")
	}
}

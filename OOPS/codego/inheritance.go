package codego

import "fmt"

/*
Inheritance is when one class (child) gets the properties and behavior of another class (parent), so we can reuse code and extend functionality.
*/
/*
Real-Time Example of Inheritance
📦 Example: Appliance → Washing Machine
Parent class: Appliance (has power ON/OFF)
Child class: WashingMachine (inherits power ON/OFF and adds Wash, Rinse, Spin)

"This way:
You don’t re-write the power functions.
You reuse and extend features in the child".
*/
/*
Go does not support inheritance in the traditional Java-style way (no extends keyword).
Instead, "Go uses "COMPOSITION" (a more flexible and preferred way in Go)."
*/

// Parent-like struct
type Animal struct{}

func (a Animal) Eat() {
	fmt.Println("This animal eats food.")
}

// Child-like struct embedding Animal
type Dog struct {
	Animal // Embedded struct (composition)
}

func (d Dog) Bark() {
	fmt.Println("Dog barks.")
}

func INHERITANCE_COMPOSITION() {
	d := Dog{}
	d.Eat()  // Inherited-like behavior
	d.Bark() // Dog's own method
}

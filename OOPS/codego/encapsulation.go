package codego

import "fmt"

/*
Encapsulation means hiding the internal details of how something works and only showing what’s necessary to the outside world.
*/

/*
Encapsulation is like putting code and data inside a box, locking it, and only giving others a key (public methods) to safely interact with it — without knowing what’s inside.
*/
/*
Real-World Analogy:
Like a TV remote – you press buttons (interface), but you don’t know how signals are sent inside (internal logic is hidden)

Real-life analogy: A Washing Machine

Imagine a washing machine:
You just press a few buttons (like Start, Wash Mode, Timer).
You don’t know the internal wiring, motor working, or how it drains and refills water.
All internal workings are hidden (encapsulated) from you.
You only interact with a clean, simple interface (buttons and display).

🔐 Encapsulation means hiding the internal details of an object and only exposing necessary parts via a public interface.
*/
/*
In Go:

Capitalized fields and methods are exported (public).
Lowercase ones are unexported (private).
Encapsulation is achieved using structs with unexported fields and getter/setter methods.
*/
// Struct with private fields (unexported)
type person struct {
	name string
	age  int
}

// Constructor-like function
func NewPerson(name string, age int) *person {
	p := &person{}
	p.SetName(name)
	p.SetAge(age)
	return p
}

// Getter and setter methods
func (p *person) GetName() string {
	return p.name
}

func (p *person) SetName(name string) {
	if len(name) > 0 {
		p.name = name
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetAge(age int) {
	if age > 0 {
		p.age = age
	}
}

func Encapsulation() {
	p := NewPerson("Bob", 30)
	fmt.Println(p.GetName()) // Bob
	fmt.Println(p.GetAge())  // 30
}

package code

import "fmt"

/*
Pass a struct as an interface to another struct’s function.
*/

// Define an interface
type Starter interface {
	Start()
}

// Struct that implements the interface
type Enginea struct{}

func (e Enginea) Start() {
	fmt.Println("Engine has started.")
}

// Another struct which uses the interface
type Cara struct{}

// Function that accepts interface as argument
func (c Cara) Drive(s Starter) {
	s.Start() // Calls the interface method
	fmt.Println("Car is now driving.")
}

func Another_structs_function() {
	engine := Enginea{} // Concrete struct
	car := Cara{}

	car.Drive(engine) // Pass struct as interface
}

package practical

import "fmt"

/*
Go (Golang) doesn’t use classes like traditional object-oriented languages (e.g., Java or C++), but you can simulate class-like behavior using structs, constructors (custom init functions), and methods (functions with receivers).

Here’s a simple example that demonstrates:

A struct (Car) representing a class
A constructor function (NewCar)
Methods using struct receivers
*/

// Car struct simulates a class
type Car struct {
	Brand string
	Model string
	Year  int
}

// Constructor function to create a new Car object
func NewCar(brand, model string, year int) Car {
	return Car{
		Brand: brand,
		Model: model,
		Year:  year,
	}
}

// Method attached to Car struct (simulating class method)
func (c Car) DisplayInfo() {
	fmt.Printf("Car: %s %s (%d)\n", c.Brand, c.Model, c.Year)
}

func Class_Object_Constructor() {
	// Creating object using constructor
	car1 := NewCar("Toyota", "Camry", 2023)

	// Calling method on object
	car1.DisplayInfo()
}

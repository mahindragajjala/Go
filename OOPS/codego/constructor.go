package codego

import "fmt"

/*
A constructor is used to insert or initialize values into a structure (struct) or object, often using user input or passed parameters, so that the object is ready to use.
*/
/*
A constructor is a "special function" used to create and initialize an object (or struct in Go). It sets up the initial state of the object when it is created.

💡 In Go, there are no built-in constructors like in other languages (Java/C++), but we can create custom constructor functions.
*/
/*
Real-Time Example
- Imagine you're building a system for a Car Rental Company.

You want to create a Car struct, and every time a car is added to the system, it must have a Brand, Model, and Year.

A constructor helps you guarantee that every Car is created properly, instead of letting someone forget to set a field.
*/

type Car_C struct {
	Brand string
	Model string
	Year  int
}

// ✅ Constructor function
func NewCar(brand, model string, year int) Car_C {
	return Car_C{Brand: brand, Model: model, Year: year}
}

// ✅ Regular method
func (c Car_C) Start() {
	fmt.Println(c.Brand, c.Model, "is starting...")
}

func GoConstructor() {
	car := NewCar("Toyota", "Fortuner", 2023) // constructor call
	car.Start()                               // function/method call
}

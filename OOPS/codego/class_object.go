package codego

import "fmt"

/*
🔹 Class:
A class is a blueprint or template for creating objects.
It defines properties (fields) and behaviors (methods) that the objects created from the class can have.

💡 Think of a class like a car design blueprint.

🔹 Object:
An object is an instance of a class.
It is created using the class blueprint and contains actual data.

💡 If the class is the blueprint, then the object is the actual car made from that blueprint.


🚗 Real-Time Example: Car Showroom

Class: Car – represents the design of a car.

Objects:
car1 – a Toyota Fortuner 2023.
car2 – a Honda City 2022.

Each car has:
Properties: brand, model, year
Behaviors: start(), stop()
*/
// Struct (like class)
type Car struct {
	Brand string
	Model string
	Year  int
}

// Method to start the car
func (c Car) Start() {
	fmt.Println(c.Brand, c.Model, "is starting...")
}

// Method to stop the car
func (c Car) Stop() {
	fmt.Println(c.Brand, c.Model, "is stopping...")
}

func ClassObject() {
	// Create objects
	car1 := Car{"Toyota", "Fortuner", 2023}
	car2 := Car{"Honda", "City", 2022}

	// Call methods
	car1.Start() // Toyota Fortuner is starting...
	car2.Stop()  // Honda City is stopping...
}

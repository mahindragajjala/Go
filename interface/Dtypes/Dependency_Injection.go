package dtypes

import "fmt"

/*
Dependency Injection means passing dependencies from the
outside rather than creating them inside a class or function.

This makes your code:
Modular
Easy to test
Easy to replace dependencies (e.g., switch database, logger, service layer).

Real-World Example
Imagine a Car object.
A car needs an Engine to run.

❌ Without Dependency Injection:
The car decides what engine it uses, hardcoded inside:

Car {
    Engine = PetrolEngine()
}
Now if you want to replace the engine with DieselEngine or ElectricEngine, you must change the Car code.

*/

//define interface
// Engine interface
type Engine interface {
	Start() string
}

//Concrete Implementations
type PetrolEngine struct{}

func (p PetrolEngine) Start() string {
	return "Petrol Engine started!"
}

type ElectricEngine struct{}

func (e ElectricEngine) Start() string {
	return "Electric Engine started!"
}

//Car Struct with Dependency Injection
type Car struct {
	engine Engine
}

func NewCar(e Engine) *Car {
	return &Car{
		engine: e,
	}
}

func (c *Car) Start() {
	fmt.Println(c.engine.Start())
}
func Dependency_injection() {
	petrol := PetrolEngine{}
	electric := ElectricEngine{}

	car1 := NewCar(petrol)
	car2 := NewCar(electric)

	car1.Start() // Output: Petrol Engine started!
	car2.Start() // Output: Electric Engine started!
}

package code

import "fmt"

type Car struct {
	engine Engine
}

func (c *Car) Start() {
	fmt.Println(c.engine.Start())
}

// Engine interface
type Engine interface {
	Start() string
}
type PetrolEngine struct{}

func (p PetrolEngine) Start() string {
	return "Petrol Engine started!"
}

type ElectricEngine struct{}

func (e ElectricEngine) Start() string {
	return "Electric Engine started!"
}

type MockEngine struct{}

func (m MockEngine) Start() string {
	return "Mock Engine started for testing."
}
func NewCar(e Engine) *Car {
	return &Car{
		engine: e,
	}
}
func Dependency_injection() {
	petrol := PetrolEngine{}
	electric := ElectricEngine{}

	car1 := NewCar(petrol)
	car2 := NewCar(electric)

	car1.Start()
	car2.Start()
}

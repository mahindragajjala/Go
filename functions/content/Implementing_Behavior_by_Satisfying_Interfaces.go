package content

import (
	"fmt"
	"math"
)

/*
Implementing Behavior by Satisfying Interfaces
In Go, an interface defines a set of method signatures.

Any type that implements all the methods in the interface automatically satisfies the interface

— no need for explicit declarations.
*/
// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Implement the interface with a struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Another implementation
type Rectangle_ struct {
	Width, Height float64
}

func (r Rectangle_) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle_) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Use interface as a parameter
func printDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n\n", s.Perimeter())
}

func Implementing_Behavior_by_Satisfying_Interfaces() {
	c := Circle{Radius: 5}
	r := Rectangle_{Width: 4, Height: 6}

	printDetails(c)
	printDetails(r)
}

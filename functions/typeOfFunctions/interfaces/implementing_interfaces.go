package interfaces

import (
	"fmt"
	"math"
)

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
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Use interface as a parameter
func printDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n\n", s.Perimeter())
}

func Interfaces_in_the_functions() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	printDetails(c)
	printDetails(r)
}

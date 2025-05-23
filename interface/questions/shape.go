package questions

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Implementing Area() for Rectangle
func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Implementing Perimeter() for Rectangle
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct
type Circle struct {
	radius float64
}

// Implementing Area() for Circle
func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implementing Perimeter() for Circle
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Function to display area and perimeter
func Interface_data(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

// Function to insert and process shapes
func Inserting_Data() {
	rectangle := &Rectangle{Width: 23, Height: 45}
	Interface_data(rectangle)

	circle := &Circle{radius: 10}
	Interface_data(circle)
}

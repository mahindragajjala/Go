package code

import (
	"fmt"
	"math"
)

/*
Polymorphism with Interfaces
Pass interfaces, not structs.
Write functions that accept interfaces, test with different structs.
in go
*/

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Write Function that Accepts Interface
func PrintArea(s Shape) {
	fmt.Printf("The area is: %.2f\n", s.Area())
}

// Use with Different Structs
func Use_with_Different_structs() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 4}
	PrintArea(c)
	PrintArea(r)
}

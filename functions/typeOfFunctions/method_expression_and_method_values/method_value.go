package methodexpressionandmethodvalues

import (
	"fmt"
	"math"
)

/*
A method value is a method bound to a specific receiver (object).
*/
type Circle struct {
	Radius float64
}

// Method with value receiver
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Both_method_by_expression_value() {
	c := Circle{Radius: 5}

	// Method Value
	areaVal := c.Area                       // Bound to c
	fmt.Println("Method Value:", areaVal()) // No args needed

	// Method Expression
	areaExpr := Circle.Area                        // Not bound to c
	fmt.Println("Method Expression:", areaExpr(c)) // c passed as arg
}

/*

Syntax:
c := Circle{Radius: 5}
areaFunc := c.Area
fmt.Println(areaFunc()) // method bound to 'c'


c.Area creates a function value that closes over (remembers) the receiver c.
When you call areaFunc(), it's equivalent to calling c.Area().
The receiver c is implicitly captured and does not need to be passed again.

Use Case:
Good when you want a method tied to a particular object and reuse it later, like callbacks or passing to goroutines.
*/

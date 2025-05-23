package coupled

import "fmt"

/*
Create two structs and make one struct use a method of the other without using interfaces.
*/
type a struct {
}

func (x *a) AAA() {
	fmt.Println("AAA is printing...")
}

type b struct {
	a *a
}

func (Y *b) BBB() {
	Y.a.AAA()
}

func A() {
	m := b{a: &a{}}
	m.BBB()
}

/*
Embed one struct inside another and access the embedded struct’s method.
*/
type x struct {
}

func (x *x) X() {
	fmt.Println("your printing the x function")
}

type y struct {
	x
}

func (y *y) Y() {
	fmt.Println("your printing the y function")
}
func (y *y) Embedded() {
	y.X()
	y.Y()
}
func Embedded_Print() {
	z := y{}
	z.Embedded()
}

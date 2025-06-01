package code

import "fmt"

type InterfaceData interface {
	Functionality()
}

type A struct {
	Data1 int
	Data2 int
}

func (a A) Functionality() {
	fmt.Println(a.Data1)
}

type B struct {
	Data1 int
	Data2 int
}

func (b B) Functionality() {
	fmt.Println(b.Data1)
}

type C struct {
	Data1 int
	Data2 int
}

func (c C) Functionality() {
	fmt.Println(c.Data1)
}

type D struct {
	Data1 int
	Data2 int
}

func (d D) Functionality() {
	fmt.Println(d.Data1)
}

func InterfaceDataFunc(i InterfaceData) {
	i.Functionality()
}

func InterfaceFunction() {
	objA := A{Data1: 1, Data2: 2}
	objB := B{Data1: 3, Data2: 4}
	objC := C{Data1: 5, Data2: 6}

	InterfaceDataFunc(objA)
	InterfaceDataFunc(objB)
	InterfaceDataFunc(objC)
}

package operatorspunctuationdata

import "fmt"

/*
Use 						Case			Example	Meaning
Struct field				p.Name			Access field in struct
Method call					p.Greet()		Call method on a struct
Package member				fmt.Println()	Use item from imported package
Auto-deref pointer			ptr.Name		Access field of a pointer struct
*/

//Access Struct Fields
type Persondata struct {
	Name string
	Age  int
}

//. (dot) operator is a multi-purpose access operator
func Access_operator() {

	//Access Fields from Pointers
	p := &Persondata{"Alice", 25}

	// Package Access
	fmt.Println(p.Name) // "Alice"
	fmt.Println(p.Age)  // "25"

	p.Greet()

}

//Access Methods
func (p Persondata) Greet() {
	fmt.Println("Hi, my name is", p.Name)
}

package code

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Struct_with_fields() {
	p1 := Person{Name: "John", Age: 30}
	fmt.Println(p1)
}

/*
Where Will You Use Structs?

Real-world Data Modeling:
Represent objects like Employee, Car, Student, Product, etc.

API Requests & Responses:
When working with JSON APIs — you map JSON data to structs.

Database Models:
ORM libraries like GORM map database tables to structs.

Organized Code:
Grouping multiple values under one name makes code cleaner and easier to maintain.
*/
//Types of struct usage

//Basic Struct with Fields
type Car struct {
	Brand string
	Year  int
}

func Anonymous_struct() {
	// Anonymous Structs
	p := struct {
		Name string
		Age  int
	}{Name: "Alice", Age: 25}

	fmt.Println(p)
}

// Nested Structs
func Nested_structs() {
	type Address struct {
		City  string
		State string
	}

	type Employee struct {
		Name    string
		Address Address
	}
	fmt.Println(Address{}, Employee{})
}

func Pointer_to_Struct() {
	p := &Person{Name: "Bob", Age: 40}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
func Data() {
	data := Person{}
	fmt.Println("---->", data.Age)
}

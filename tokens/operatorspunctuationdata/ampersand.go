package operatorspunctuationdata

import "fmt"

//&
//Passing by Reference to Functions
func Passing_by_Reference_to_Functions() {
	x := 5
	update(&x)
	fmt.Println("Pass by Pointer:", x)

	y := 5
	checking(y)
	fmt.Println("Pass by Value:", y)

}

//Pass by Pointer (&x)
func update(val *int) {
	*val = *val + 10
}

//Pass by Value (x)
func checking(y int) {
	y = y + 5
	fmt.Println(y)
}

//Working with Structs (Memory-efficient passing)
//Instead of copying the entire struct, you can pass a pointer to the struct.

type Person struct {
	name string
	age  int
}

func birthday(p *Person) {
	p.age++
}

func Working_with_Structs_of_ampersand() {
	p := Person{name: "Mahi", age: 24} //object
	birthday(&p)
	fmt.Println(p.age)
}

//Using new() to Allocate Memory
/*
Go’s new() function returns a pointer, and we use & when we want the
address of a value we already created.
*/

func Ampersand_with_New() {
	x := 100
	p := &x // address of x
	q := new(int)
	fmt.Printf("New keyword type :%T \n new data : %v\n", q, q)
	*q = 200
	fmt.Println(*p, *q) // Output: 100 200
}

//Linking Data Structures (like Linked Lists, Trees, etc.)
//Pointers are necessary to build complex data structures.

type Node struct {
	value int
	next  *Node
}

func Ampersand_Linking_Data_Structures() {
	Nodedata := Node{value: 25}
	fmt.Printf("%T", Nodedata.next)
	fmt.Printf("%T", Nodedata.value)

}

// Efficiently Updating "Slice or Map" Fields in Structs

type Data struct {
	values []int
}

func (d *Data) AddValue(v int) {
	d.values = append(d.values, v)
}

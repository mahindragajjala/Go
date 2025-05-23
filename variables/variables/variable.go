package variables

import "fmt"

/*

A variable is a named place in memory where a value is stored.
You can use variables to store, retrieve, and manipulate data in your program.

--> Types of variables
Every variable has a type. The type defines what kind of data the variable can hold:
int for integers
string for text
bool for true/false
interface{} for any type (more on this later)

---> Declaring Variables
You can declare variables using the var keyword.
When you declare a variable, it reserves space in memory.

---> Structured variables
Some variables are more complex:
Arrays, slices, and structs are made of elements or fields.
Each field or element is also like its own variable.

----> Static vs Dynamic Type
* Static type is the type given when the variable is declared.
* Dynamic type applies to interface variables, and it’s the type of
the actual value assigned at runtime.

----> Interface Example
var x interface{} // static type: interface{}, value: nil
var v *T          // static type: *T, value: nil
x = 42            // value: 42, dynamic type: int
x = v             // value: (*T)(nil), dynamic type: *T
Even if x is of type interface{}, what it actually holds (at a given moment) determines its dynamic type.

---->Default (Zero) Values
If a variable is declared but not assigned a value, it gets a default value:
For int: 0
For string: ""
For pointers: nil

Concept	                       Description
Variable	                   A place to store a value
Static Type	                   Type declared at compile-time
Dynamic Type	               Actual type of value held (for interface variables)
Zero Value	                   Default value if nothing is assigned
Structured Types	           Arrays, slices, structs with individual elements/fields
Pointers	                   Used to refer to dynamically created variables

*/
//types of variables
var age int     // age can hold an integer
var name string // name can hold text

//Declaring Variables
var x int

// Structured Variables
type Person struct {
	name string
	age  int
}

func Variable() {
	insertion := Person{name: "mahi", age: 23}
	fmt.Println(insertion)

	var xy interface{} // x can hold any value, but it starts out as nil
	xy = 42            // now x has dynamic type int
	xy = "hello"       // now x has dynamic type string
	fmt.Println(xy)
	fmt.Println(age, name, x)
}

package typekeyword

/*
Use Case					Example Syntax
Struct	                    type MyStruct struct { ... }
Interface	 				type MyInterface interface { ... }
Named type	 				type Celsius float64
Type alias					type MyFloat = float64
Function type				type Operation func(int, int) int
*/
//Creating a New Struct
type TypeStructure struct {
	A int
	b string
}

//Defining a New Interface
type TypeStructure_interface interface {
	Typekeyword()
}

//Creating Type Aliases

//Named type (not interchangeable with original type without conversion)
type Age int

//Type alias (interchangeable with original type)
type MyInt = int

//Function Type Definitions
type Adder func(int, int) int

//Custom Error Types or Wrapper Types
type MyError struct {
	Msg string
}

func (e MyError) Error() string {
	return e.Msg
}

func Typekeyword() {

}

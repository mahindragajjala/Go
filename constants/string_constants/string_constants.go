package stringconstants

import "fmt"

/*
Feature	                 Example
Interpreted string       "Hello\nWorld"
Raw string	             `Hello\nWorld`
Typed constant	         const s string = "hi"
Untyped constant	     const s = "hi"
Concatenation	         const c = a + b
Indexing	             s[0] gives byte
A string constant in Go:

Is a read-only sequence of bytes (UTF-8 encoded)

Can be typed or untyped

Is immutable (cannot be changed once defined)
*/
func StringConstants() {
	//Declaring String Constants
	const message = "Hello, Go!" // untyped string constant
	const typedMessage string = "Typed string"

	//Types of String Literals

	//1. Interpreted String Literals (double quotes "")
	//Supports escape sequences like \n, \t, \\, etc.
	const msg = "Hello,\nWorld!"

	//Raw String Literals (backticks ` `)
	//Multiline allowed, no escape sequences
	const raw = `Line 1
	Line 2\t <- not escaped
	Line 3`

	fmt.Println(message, typedMessage, msg, raw)
	//Concatenation of String Constant
}

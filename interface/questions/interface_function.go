package questions

import "fmt"

/*
21. **Use interface values to store different types** (`int`, `string`, `struct`) and print them.
*/
func Interface_function(v interface{}) {
	fmt.Println(v)
}
func Interface_main() {
	Interface_function("inputdata")
	Interface_function(23)
	Interface_function(23.99)
}

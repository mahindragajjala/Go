package pointer

import "fmt"

//Functions can accept pointers to modify the original value.

func Pointer_with_function() {
	var a int = 10
	var b int = 20
	fmt.Println("BEFORE:", a, b)
	Reference_function(&a, &b)
	fmt.Println("AFTER:", a, b)

}
func Reference_function(a *int, b *int) {
	*a = *a + 10
	*b = *b + 10
}

package operatorspunctuationdata

import "fmt"

/*

The , (comma) operator in Go isn't an
operator in the same way it is in C or C++.
In Go, it's more like a separator

*/
func Separator() {
	//Multiple variable declarations
	a, b := 1, 2

	//Function arguments and return values
	fmt.Println(a, b)

	//Multiple return values from functions

	j, k := Another_separator(1, 2)
	fmt.Println(j, k)

	//Ignoring values using _ (blank identifier)
	slicedata := []int{1, 2, 3, 4, 5, 6}
	for _, data := range slicedata {
		fmt.Println(data)
	}

	//Multiple values in range
	for i, v := range []string{"go", "lang"} {
		fmt.Println(i, v)
	}

	//Composite literals
	//Values in arrays, slices, maps, and structs use comma as a separator.
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

}
func Another_separator(a int, b int) (int, int) {
	return 3, 4
}

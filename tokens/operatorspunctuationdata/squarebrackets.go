package operatorspunctuationdata

import "fmt"

func Square_brackets() {
	//Arrays and Slices
	//Array and Slice Declaration
	var arr [5]int
	Array := [2]int{1, 2}
	slice := []int{1, 2, 3}
	fmt.Println(Array, slice)
	fmt.Printf("\n%T", arr)

	//Slicing
	s := []int{1, 2, 3, 4, 5}
	sub := s[1:4]    // elements at index 1, 2, 3
	fmt.Println(sub) // Output: [2 3 4]

	//Maps
	mapdata := map[string]int{"one": 1, "two": 2}
	fmt.Println(mapdata)

	//Type Specifier
	//[...] when creating arrays without specifying the size
	arrd := [...]int{10, 20, 30} // Compiler infers the size
	fmt.Println(arrd)
}

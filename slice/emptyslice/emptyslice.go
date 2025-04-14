package emptyslice

import "fmt"

func EmptySlice() {
	//MAKE keyword - This creates a new hidden array behind the scenes.
	EmptySlice := make([]int, 23, 40)
	fmt.Println("Actally length", len(EmptySlice))
	fmt.Println("Actally capacity", cap(EmptySlice))

	//APPEND keyword
	EmptySlice = append(EmptySlice, 45)
	fmt.Println("after adding the len", len(EmptySlice))
	fmt.Println("after adding the cap", cap(EmptySlice))

	var Empty []int
	//slice  literal in go
	/* var Empty []int
	Declares a nil slice.
	No memory is allocated for the underlying array.
	len(Empty) is 0.
	cap(Empty) is 0.
	Empty == nil is true.
	So it’s an uninitialized slice pointer pointing to nil. */

	empty := []int{}
	/* empty := []int{}
	Declares an empty but non-nil slice.
	Memory is allocated for the slice header, but the array part is zero-length.
	len(empty) is 0.
	cap(empty) is 0.
	empty == nil is false.
	So this is initialized, just empty. */

	fmt.Println("slice  literal length", len(Empty))
	fmt.Println("slice  literal capacity", cap(Empty))
	fmt.Println("type:", Empty)

	fmt.Println("var Empty []int:", Empty == nil)
	fmt.Println("empty := []int{}", empty == nil)

}

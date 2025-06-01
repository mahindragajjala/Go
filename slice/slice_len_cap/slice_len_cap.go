package slicelencap

import "fmt"

//How do you use the remaining capacity
func Filling_Remaining_Capacity() {
	//Initial slice
	EmptySlice := make([]int, 23, 40)
	fmt.Println(len(EmptySlice))
	fmt.Println(cap(EmptySlice))

	// Adding elements up to the capacity
	for i := len(EmptySlice); i < cap(EmptySlice); i++ {
		EmptySlice = append(EmptySlice, i)
	}

	fmt.Println("After Appending:")
	fmt.Println("Length:", len(EmptySlice))
	fmt.Println("Capacity:", cap(EmptySlice))
	fmt.Println("Slice Data:", EmptySlice)
}

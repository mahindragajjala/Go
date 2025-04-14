package slicesofslices

import "fmt"

func SlicesOfSlices() {
	SlicesOfSlices := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(SlicesOfSlices)

	//Accessing the elements
	fmt.Println("slices of slices : ", SlicesOfSlices[1][2])

	//modify elements
	SlicesOfSlices[0][1] = 5

	//Append to Inner Slice

	SlicesOfSlices[1] = append(SlicesOfSlices[1], 7, 8, 9, 10)
	fmt.Println(SlicesOfSlices)

	//Append a New Slice
	SlicesOfSlices = append(SlicesOfSlices, []int{123, 456, 789})
	fmt.Println(SlicesOfSlices)

	//Delete a slice
	SlicesOfSlices = append(SlicesOfSlices[:1], SlicesOfSlices[2:]...)
	fmt.Println(SlicesOfSlices) // Output: [[1 5 3] [123 456 789]]

	//Iterating over elements
	for i, inner := range SlicesOfSlices {
		fmt.Printf("Row %d: %v\n", i, inner)
	}

	// Length of Slices
	fmt.Println("Outer slice length:", len(SlicesOfSlices))   // Outer slice length
	fmt.Println("Inner slice length", len(SlicesOfSlices[0])) // Inner slice length

	//Copy Inner Slice
	copySlice := make([]int, len(SlicesOfSlices[0]))
	copy(copySlice, SlicesOfSlices[0])
	fmt.Println("Copy Inner Slice:", copySlice) // Output: [1 2 3 7]

	//Re-slicing Inner Slice
	sub := SlicesOfSlices[0][1:3]
	fmt.Println(sub) // Output: [2 3]

}

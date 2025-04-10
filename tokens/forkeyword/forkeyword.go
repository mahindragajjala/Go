package forkeyword

import "fmt"

func For_keyword() {
	//Basic for Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for as a while Loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//Infinite Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Looping Over Arrays, Slices, Maps (using range)
	nums := []int{10, 20, 30}
	for i, val := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, val)
	}

}

package pointer

import (
	"fmt"
	"time"
)

// Large structure with 1 million integers
type LargeData struct {
	data [1000000]int
}

// Function that accepts LargeData by value
func processByValue(ld LargeData) {
	ld.data[0] = 999
}

// Function that accepts LargeData by reference (pointer)
func processByPointer(ld *LargeData) {
	ld.data[0] = 999
}

func Large_Data() {
	// Create large data
	bigData := LargeData{}

	start := time.Now()
	processByValue(bigData)
	fmt.Println("Time taken by passing by value:", time.Since(start))

	start = time.Now()
	processByPointer(&bigData)
	fmt.Println("Time taken by passing by pointer:", time.Since(start))
}

package questions

import "fmt"

/*
Write a Go program to reverse a map — swap keys and values. (Assume values are unique.)
*/
func reverseMap(original map[string]int) map[int]string {
	reversed := make(map[int]string)
	for key, value := range original {
		reversed[value] = key
	}
	return reversed
}

func Reversed_string() {
	// Original map
	originalMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Reverse the map
	reversedMap := reverseMap(originalMap)

	// Print the reversed map
	fmt.Println("Original Map:", originalMap)
	fmt.Println("Reversed Map:", reversedMap)
}

package questions

import "fmt"

/*
Write a program to count the frequency of each element in an array using a for loop.
Let the array be: [1, 2, 2, 3, 3, 3, 4]

➡️ Expected Output:
1 occurs 1 time
2 occurs 2 times
3 occurs 3 times
4 occurs 1 time
➡️ Goal: Combine loop logic with maps/dictionaries (or arrays, depending on language) to track counts.
*/

func Count_the_frequency_of_element() {
	// Input array
	arr := []int{1, 2, 2, 3, 3, 3, 4}

	// Step 1: Create a map to store frequencies
	frequency := make(map[int]int)

	// Step 2: Loop through the array to count frequencies
	for _, val := range arr {
		frequency[val]++ // Increment count for this value
	}

	// Step 3: Print each element and its count
	for key, count := range frequency {
		fmt.Printf("%d occurs %d time", key, count)
		if count > 1 {
			fmt.Print("s")
		}
		fmt.Println()
	}
}

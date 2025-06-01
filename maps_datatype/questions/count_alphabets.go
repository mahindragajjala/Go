package questions

import (
	"fmt"
)

/*
Write a Go program to count the frequency of each character in a string using a map.
*/
func Count_the_frequency_of_each_character_in_a_string() {
	// Input string
	input := "hello world"

	// Create a map to store character frequencies
	charFrequency := make(map[rune]int)

	// Loop through each character in the string
	for _, char := range input {
		fmt.Println(char)
		charFrequency[char]++
	}

	// Print the character frequencies
	fmt.Println("Character frequencies:")
	for char, count := range charFrequency {
		fmt.Printf("'%c' : %d\n", char, count)
	}
}

/*
Step	Character	Map Status
1	'h'	{ 'h': 1 }
2	'e'	{ 'h': 1, 'e': 1 }
3	'l'	{ 'h': 1, 'e': 1, 'l': 1 }
4	'l'	{ 'h': 1, 'e': 1, 'l': 2 }
5	'o'	{ 'h': 1, 'e': 1, 'l': 2, 'o': 1 }
6	' '	{ 'h': 1, 'e': 1, 'l': 2, 'o': 1, ' ': 1 }
7	'w'	{ ..., 'w': 1 }
8	'o'	{ ..., 'o': 2 }
9	'r'	{ ..., 'r': 1 }
10	'l'	{ ..., 'l': 3 }
11	'd'	{ ..., 'd': 1 }
*/

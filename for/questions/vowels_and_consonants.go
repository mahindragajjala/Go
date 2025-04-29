package questions

import "fmt"

/*
Write a program to count the number of vowels and consonants in a string using a for loop.
Let the string be: "programming"

➡️ Expected Output:

Vowels: 3

Consonants: 8

➡️ Goal: Work with character checking inside a loop.
*/
func Vowels_and_consonants() {
	input := "programming"
	vowels := "aeiouAEIOU" // Include uppercase too if needed
	vowelCount := 0
	consonantCount := 0

	for j := 0; j < len(input); j++ {
		char := input[j]
		isVowel := false

		// Check if current char is vowel
		for i := 0; i < len(vowels); i++ {
			if char == vowels[i] {
				isVowel = true
				break
			}
		}

		if isVowel {
			vowelCount++
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// Consonant check (must be an alphabet but not a vowel)
			consonantCount++
		}
	}

	fmt.Println("Vowels:", vowelCount)
	fmt.Println("Consonants:", consonantCount)
}

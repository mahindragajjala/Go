package questions

import "fmt"

func Remove_Vowels_and_consonants() {
	input := "programming"
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	vowelCount := 0
	consonantCount := 0

	for j := 0; j < len(input); j++ {
		char := rune(input[j])

		if vowels[char] {
			vowelCount++
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// Consonant check (must be an alphabet but not a vowel)
			consonantCount++
		}
	}

	fmt.Println("Vowels:", vowelCount)
	fmt.Println("Consonants:", consonantCount)
}

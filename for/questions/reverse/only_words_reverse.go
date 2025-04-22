package questions

import (
	"fmt"
	"strings"
)

func Only_Words_reverse() {

	string := "Hey siri"
	// Split the sentence into words
	words := strings.Fields(string)

	// Use for loop to reverse the slice
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the words back into a string
	data := strings.Join(words, " ")
	fmt.Println(data)
}

package stringtypes

import "fmt"

func StringAreImmutable() {
	//but we can change data in it:) using byte and rune
	MutableTOImmutable_Byte()
	MutableTOImmutable_Rune()
}

func MutableTOImmutable_Byte() {
	s := "Hello"
	// slice of bytes.
	b := []byte(s)  // Convert string to byte slice
	b[0] = 'M'      // Modify the first byte
	b[1] = 'M'      // Modify the first byte
	b[2] = 'M'      // Modify the first byte
	b[3] = 'M'      // Modify the first byte
	b[4] = 'M'      // Modify the first byte
	s2 := string(b) // Convert back to string

	fmt.Println("Original:", s)  // Output: Hello
	fmt.Println("Modified:", s2) // Output: Mello
}

func MutableTOImmutable_Rune() {
	s := "नमस्ते"
	r := []rune(s) // Convert to rune slice (handles Unicode properly)

	r[0] = 'श'      // Replace the first character
	s2 := string(r) // Convert back to string

	fmt.Println("Original:", s)  // Output: नमस्ते
	fmt.Println("Modified:", s2) // Output: शमस्ते
}

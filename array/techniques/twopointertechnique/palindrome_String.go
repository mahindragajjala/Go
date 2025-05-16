package twopointertechnique

import "fmt"

/*
3. Check if String is Palindrome
   *Use two pointers to check if a string reads the same forwards and backwards.*
*/

func Palindrome_string(data string) bool {
	fmt.Println(len(data))
	left := 0
	right := len(data) - 1

	for left < right {
		if data[left] != data[right] {
			return false
		}
		left++
		right--
	}

	return true
}

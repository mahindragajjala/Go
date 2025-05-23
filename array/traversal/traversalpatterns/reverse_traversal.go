package traversalpatterns

import (
	"fmt"
	"strings"
)

// Print all elements of an array in reverse order.
var reversearray = []int{1, 2, 3, 4, 5, 6}

func Reverse_order() {

	for i := len(reversearray) - 1; i >= 0; i-- {
		fmt.Println(reversearray[i])
	}
}

// Reverse a string manually using loop.
func ReverseString() {
	data := "BANANA"
	for i := len(data) - 1; i >= 0; i-- {
		fmt.Print(string(data[i]))
	}
}

// Reverse an array in-place.
func Reverse_in_place() {
	data := []rune("BANANA") // Convert string to mutable rune slice

	length := len(data)
	for i := 0; i < length; i++ {
		for j := length - 1; j >= 0; j-- {
			if i < j {
				data[i], data[j] = data[j], data[i]
				break // After one swap for each i, break the inner loop
			}
		}
		if i >= length/2 {
			break // Stop when the middle is reached
		}
	}

	fmt.Println("Reversed:", string(data))
}
func Reverse_in_place_banana() {
	data := []rune("BANANA") // Mutable slice

	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	fmt.Println("Reversed:", string(data))
}

// Find the last occurrence of a target element.
func Last_occurence_reverse() {
	arr := []int{5, 3, 7, 3, 9, 3}
	target := 3
	var occurrence int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			occurrence = i
		}
	}
	fmt.Println(occurrence)
}

//Remove trailing zeroes from an integer array.

// Removing anta...
func Removing_zeroes() {
	arr := []int{4, 5, 0, 0}
	var newarray []int
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			newarray = append(newarray, arr[i])
		}
	}
	fmt.Println(newarray)
}

func Trailing_zeroes() {
	arr := []int{4, 0, 5, 0, 0}
	end := len(arr) - 1

	// Move backward to find the last non-zero element
	for end >= 0 && arr[end] == 0 {
		end--
	}

	// Slice the array up to the last non-zero element
	newArray := arr[:end+1]
	fmt.Println(newArray)
}

/* 2. Reverse a string manually using loop.
6. Print string characters in reverse order.
7. Reverse words in a sentence (without using built-in functions). */

// input := "Go is great"
// manually
func ReverseWords_manually(sentence string) {
	n := len(sentence)
	word := ""   //Creates an empty string to build each word in reverse.
	result := "" //Creates an empty string to hold the final reversed sentence.

	for i := n - 1; i >= 0; i-- {
		if sentence[i] == ' ' { //Checks if the current character is a space (indicating the end of a word when reading backwards).
			if word != "" {
				result += word + " "
				word = ""
			}
			/*
				If a word has been built, it's added to the result followed by a space.
				Then word is cleared to start building the next one.
			*/
		} else {
			// Add character to the front of the word
			word = string(sentence[i]) + word
			/*
				If not a space, keep building the word by prepending the character.
				This ensures the word is built in the correct order, even though you're moving backward.
			*/
		}
	}
	if word != "" {
		result += word
	}

	fmt.Println(result)
}

/*
| i  | sentence\[i] | word  | result      |
| -- | ------------ | ----- | ----------- |
| 10 | t            | t     |             |
| 9  | a            | at    |             |
| 8  | e            | eat   |             |
| 7  | r            | reat  |             |
| 6  | g            | great |             |
| 5  | (space)      |       | `great `    |
| 4  | s            | s     | `great `    |
| 3  | i            | is    | `great `    |
| 2  | (space)      |       | `great is ` |
| 1  | o            | o     | `great is ` |
| 0  | G            | Go    | `great is ` |
*/
func ReverseWords_keywords(sentence string) {
	words := strings.Fields(sentence) // Splits by space, trims extras
	n := len(words)

	// Reverse the words slice
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}

	// Join back to form the reversed sentence
	result := strings.Join(words, " ")
	fmt.Println(result)
}

/*
| Step    | Value/Operation | Result                  |
| ------- | --------------- | ----------------------- |
| Input   | "Go is great"   |                         |
| Fields  | Split words     | `["Go", "is", "great"]` |
| Reverse | Swap 0↔2        | `["great", "is", "Go"]` |
| Join    | Make string     | `"great is Go"`         |
| Output  | Print           | `great is Go`           |

*/

//Detect palindrome using reverse traversal only.

func Detect_Palidrome() {
	array := "madam"
	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j >= 0; j-- {
			if array[i] == array[j] {
				fmt.Println("palindrome")
			} else {
				fmt.Println("not palindrome")
			}
		}
	}
}
func Detect_Palindrome() {
	array := "madam"
	length := len(array)
	isPalindrome := true

	// Only reverse traversal on the right side (j)
	// Compare array[i] with array[j] where j = length-1 - i
	for i := 0; i < length; i++ {
		j := length - 1 - i // reverse index
		if array[i] != array[j] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}

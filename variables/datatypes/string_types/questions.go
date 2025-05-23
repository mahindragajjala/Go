package stringtypes

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var DeclarationVariable string = "Declaration & insert data"
var DeclarationVariable2 string = "Declaration & insert data"
var DeclarationVariable3 string = "Declaration"

func Declaration() {
	/*
	   Print a String:
	   Write a Go program to declare a string and print it.
	*/
	fmt.Println("Declation:", DeclarationVariable)
}
func FindLength() {
	/*
	   String Length:
	   Write a program that takes a string and prints its length using len().
	*/
	Length := len(DeclarationVariable)
	fmt.Println(Length)
}
func Concatenate_string() {
	/*
	   Concatenate Strings:
	   Write a program to concatenate two strings using + operator.
	*/
	Concatenate_string := DeclarationVariable + "Adding"
	fmt.Println(Concatenate_string)
}

func Access_character() {
	/*
	   Access Character:
	   Write a program to print each character of a string using a loop.
	*/
	for index, char := range DeclarationVariable {
		fmt.Printf("Index: %d, Character: %c CharacterType: %T Address:%v\n", index, char, char, &char)
	}
	/*
		Index: 0, Character: D CharacterType: int32 Address:0xc00000a0e8
		Index: 1, Character: e CharacterType: int32 Address:0xc00000a108
		Index: 2, Character: c CharacterType: int32 Address:0xc00000a10c
		Index: 3, Character: l CharacterType: int32 Address:0xc00000a120
		Index: 4, Character: a CharacterType: int32 Address:0xc00000a124
		Index: 5, Character: r CharacterType: int32 Address:0xc00000a128
		Index: 6, Character: a CharacterType: int32 Address:0xc00000a12c
		Index: 7, Character: t CharacterType: int32 Address:0xc00000a130
		Index: 8, Character: i CharacterType: int32 Address:0xc00000a134
		Index: 9, Character: o CharacterType: int32 Address:0xc00000a138
		Index: 10, Character: n CharacterType: int32 Address:0xc00000a13c
		Index: 11, Character:   CharacterType: int32 Address:0xc00000a140
		Index: 12, Character: & CharacterType: int32 Address:0xc00000a144
		Index: 13, Character:   CharacterType: int32 Address:0xc00000a148
		Index: 14, Character: i CharacterType: int32 Address:0xc00000a14c
		Index: 15, Character: n CharacterType: int32 Address:0xc00000a150
		Index: 16, Character: s CharacterType: int32 Address:0xc00000a154
		Index: 17, Character: e CharacterType: int32 Address:0xc00000a158
		Index: 18, Character: r CharacterType: int32 Address:0xc00000a15c
		Index: 19, Character: t CharacterType: int32 Address:0xc00000a160
		Index: 20, Character:   CharacterType: int32 Address:0xc00000a164
		Index: 21, Character: d CharacterType: int32 Address:0xc00000a168
		Index: 22, Character: a CharacterType: int32 Address:0xc00000a16c
		Index: 23, Character: t CharacterType: int32 Address:0xc00000a170
		Index: 24, Character: a CharacterType: int32 Address:0xc00000a174
	*/
}
func String_comparsion() bool {
	/*
	   String Comparison:
	   Write a function to compare two strings and return true if they are equal.
	*/
	if DeclarationVariable == DeclarationVariable2 && DeclarationVariable == DeclarationVariable3 {
		return true
	}

	return false
}
func Reverse_String() {
	/*
		Reverse a String:
		Write a program to reverse a string without using any built-in function.
	*/
	data := DeclarationVariable
	runes := []rune(data) // convert string to rune slice to handle Unicode properly

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// swapping runes
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversed := string(runes) // convert rune slice back to string
	fmt.Println("Original String:", data)
	fmt.Println("Reversed String:", reversed)
}

func Count_Vowels_and_Consonants() {
	/*
		Count Vowels and Consonants:
		Write a Go program to count the number of vowels and consonants in a string.
		Stringdata "Declaration & insert data" in global variable declared
	*/
	Stringdata := DeclarationVariable
	var vowel int = 0
	var consonants int = 0

	for _, value := range Stringdata {
		fmt.Println("value--->", value)
		if unicode.IsLetter(value) {
			switch unicode.ToLower(value) { // convert into the capital letter to small letter .
			case 'a', 'e', 'i', 'o', 'u':
				vowel++
			default:
				consonants++
			}
		}
	}

	fmt.Printf("Vowels: %v\nConsonants: %v\n", vowel, consonants)
}

func Palindrome() bool {
	/*
		Check Palindrome:
		Write a function to check if a string is a palindrome.
	*/
	Stringdata := DeclarationVariable
	leng := len(Stringdata)

	fmt.Println(Stringdata[0])
	fmt.Println(Stringdata[23])

	for i := 0; i < leng/2; i++ {
		if Stringdata[i] != Stringdata[leng-1-i] {
			return false
		}
	}
	return true
}

func Substring_Search() {

	/*
	   Substring Search:
	   Implement a function to check if a string contains a substring.
	*/

	mainStr := "Declaration & insert data"
	subStr := "insert"

	// Check if subStr exists in mainStr
	if strings.Contains(mainStr, subStr) {
		fmt.Printf("Substring '%s' found in '%s'\n", subStr, mainStr)
	} else {
		fmt.Printf("Substring '%s' NOT found in '%s'\n", subStr, mainStr)
	}

}

// STRING CONVERSION
// Converting Strings to Booleans
func String_to_Boolean() {
	str := "true"
	value, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Boolean:", value)
	}
}

// Converting Strings to Floats
func String_to_Float() {
	str := "3.14"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Float:", num)
	}
}

// Converting Strings to Integers
func String_to_Integers() {
	str := "42"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Integer:", num)
	}
}

//STRING FORMATTING
//fmt.Sprintf

func Sprintf_formatting() {
	name := "John"
	age := 30
	formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(formatted)
}

//fmt.Printf

func Printf_formatting() {
	name := "John"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

/*
Contains(): Determines whether one string contains another.
ToLower(): Converts a string to lowercase.
ToUpper(): Converts a string to uppercase.
TrimSpace(): Removes leading and trailing whitespace from a string.
Split(): Splits a string into a slice of substrings based on a specified delimiter.
*/
// String Manipulation
// Concatenating Strings
func Concatenating_String() {
	str1 := "Hello"
	str2 := "world"
	concatenated := str1 + ", " + str2 + "!"
	fmt.Println(concatenated)

	strs := []string{"Go", "is", "awesome!"}
	joined := strings.Join(strs, " ")
	fmt.Println(joined)
}

// finding the substrings
func Finding_substring() {
	str := "Golang is amazing!"
	fmt.Println(strings.Contains(str, "amazing"))
	fmt.Println(strings.Index(str, "amazing"))
}

// Replace substring
func Replace_substring() {
	str := "Golang is amazing!"
	replaced := strings.Replace(str, "amazing", "awesome", 1)
	fmt.Println(replaced)
}

// Splitting Strings
func Splitting_Strings() {
	str := "apple,banana,cherry"
	fruits := strings.Split(str, ",")
	fmt.Println(fruits)
}

// Trimming Strings
func Trimming_Strings() {
	str := "   Golang is fun!   "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str2 := "!!!Golang is powerful!!!"
	trimmed2 := strings.Trim(str2, "!")
	fmt.Println(trimmed2)
}

/*

Word Count:
Write a function that counts the number of words in a string.

Replace Substring:
Write a program that replaces all occurrences of one substring with another.

String Split:
Write a program to split a string by spaces and print each word.

Case Conversion:
Write a program to convert a string to upper case and lower case.

Remove Whitespace:
Write a function to trim leading and trailing spaces from a string.

Unicode Character Count:
Write a program to count the number of Unicode characters (runes) in a string.

String to Rune Slice and Back:
Write a function to convert a string to a slice of runes and back to a string.

Anagram Checker:
Write a function to check if two strings are anagrams.

Longest Substring Without Repeating Characters:
Write a program to find the length of the longest substring without repeating characters.

Permutations of a String:
Write a recursive program to print all permutations of a string.

Frequency of Characters:
Write a program that counts the frequency of each character in a string and prints the result as a map.

Custom Split Function:
Write your own function to split a string by any delimiter (don’t use strings.Split).

String Compression:
Write a function that performs basic string compression using the counts of repeated characters.
E.g. aabcccccaaa becomes a2b1c5a3.

Implement Substring Search (KMP Algorithm):
Write a function to search for a substring using the KMP algorithm.

Multibyte Safe Reversing:
Write a function that reverses a string while correctly handling multibyte characters (runes).

Immutable String Transformation:
Write a function to transform a string to title case without modifying the original string.

Concurrent String Processor:
Design a program where a string is sent through a channel, converted to upper case in one goroutine, and the result is printed in another.

Escape and Unescape JSON String:
Write a function to escape and unescape special characters in a JSON string.

Build a Simple Template Engine:
Create a function that replaces placeholders in a string with given values.
E.g. "Hello, {{name}}" → "Hello, John".

Decode UTF-8 Byte Array:
Write a program to decode a byte array into a valid UTF-8 string.

Simulate String Interning:
Design a simple string interning system in Go for memory optimization.

Bonus Conceptual:

Why are strings immutable in Go?
How does Go handle Unicode under the hood?
When to use rune vs byte when dealing with strings?
Differences between strings.Builder and concatenation.


*/

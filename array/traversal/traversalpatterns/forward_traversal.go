package traversalpatterns

import "fmt"

// FORWARD TRAVERSAL
var array = []int{2, 5, 24, 787}

func Forward_Traversal() {

	for _, value := range array {
		fmt.Println(value)
	}
}

//Searching

func Searching() {
	var target int = 24
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			fmt.Println("found...", i)
		} else {
			fmt.Println("not found...")
		}
	}
}

//Applying transformations or filters

/*
In Go (Golang), applying transformations or filters on an array means processing the array's elements to either:

Transform them (change each element into something else), or

Filter them (keep only certain elements based on a condition).
*/
/*
Transformation in an Array
This means changing every element in an array to something else (like multiplying each number by 2, converting strings to uppercase, etc.).
*/
func Transformation() {
	for i := 0; i < len(array); i++ {
		array[i] = array[i] * 2
	}
}

/*
Filtering means selecting only the elements that satisfy a condition (like all even numbers, all elements > 10, etc.).
*/

func Filtering() {
	var filter []int
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			filter = append(filter, array[i])
		} else {
			fmt.Println("odd number :", array[i])
		}
	}
	fmt.Println(filter)
}

//REVERSE TRAVERSAL

func Reverse_traversal() {
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Println(array[i])
	}
}

// Reverse a string
var array_string = []string{"a", "b", "c", "d"}

func Reverse_traversal_String() {
	for i := len(array_string) - 1; i >= 0; i-- {
		fmt.Println(array_string[i])
	}
}

//Processing Suffixes
/*
A suffix is a substring that starts from any position in a string and goes to the end of the string.
Processing suffixes means performing operations on all suffixes of a string (or sometimes an array). It is a common task in pattern matching, string algorithms, and dynamic programming.

Suffix arrays: Used in string pattern matching, search engines.

Suffix trees: Help solve problems like longest common substring, pattern occurrence.

Dynamic programming: Sometimes you need to process suffixes to build up a solution from the end (e.g., suffix sums, suffix products).

Text compression: Algorithms like Burrows-Wheeler Transform use suffix arrays.

*/

func Processing_suffixes() {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i:])
	}
}

//Finding the Last Occurrence of an Element
/*
This refers to finding the last (rightmost) index where a specific element appears in a string or array.

In string searching: find the last occurrence of a character.
In log processing: find the last time an event happened.
In arrays: find the last index where a value appears.
*/
func lastOccurrence(arr []int, target int) int {
	lastIndex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			lastIndex = i
		}
	}
	return lastIndex
}

func Last_occurence() {
	arr := []int{1, 3, 5, 3, 9, 3}
	target := 3
	fmt.Printf("Last occurrence of %d is at index %d\n", target, lastOccurrence(arr, target))
}

//CODING QUESTIONS

// Print all elements of an array.
var arr = []int{4, 2, 1, 6, 8, 3, 6}

func Printing_All() {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
}

//Find the sum of all elements in an array.

func Sum_of_elements() {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
}

// Count the number of even and odd elements.
func Count_even_odd() {
	var even int = 0
	var odd int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Printf("EVEN:%v |ODD:%v", even, odd)
}

//Find the maximum and minimum in an array.

func Max_Mini() {
	var Max int = 0
	var Mini int = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < Mini {
			Mini = arr[i]
		} else if arr[i] > Max {
			Max = arr[i]
		} else {

		}
	}
}

//Linear search for a given element.

/*
Linear Search (also called sequential search) is one of the simplest searching algorithms.

It is used to find the position of a target value (called a key) within a list, array, or collection by checking each element one by one.
*/
func Linear_Search() {
	target := 5
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			fmt.Println(i)
		}
	}
}

// Print characters of a string one by one.
var stringdata string = "hello"

func Print_Characters() {
	for i := 0; i < len(stringdata); i++ {
		fmt.Println(string(stringdata[i]))
	}
}

// Count the number of vowels in a string.
func Finding_Vowels() {
	var vowels int = 0
	var consonants int = 0
	for i := 0; i < len(stringdata); i++ {
		if string(stringdata[i]) == "A" || string(stringdata[i]) == "E" || string(stringdata[i]) == "I" || string(stringdata[i]) == "O" || string(stringdata[i]) == "U" || string(stringdata[i]) == "a" || string(stringdata[i]) == "e" || string(stringdata[i]) == "i" || string(stringdata[i]) == "o" || string(stringdata[i]) == "u" {
			vowels++
		} else {
			consonants++
		}
	}
	fmt.Println("Vowels:", vowels)
	fmt.Println("Consonants:", consonants)
}

/*
Check if a string is a palindrome using forward traversal (build a reversed string separately).
*/
var Palindrome_Data string = "aabbaa"

func Palindrome_checking() string {
	reversed := ""
	for i := len(Palindrome_Data) - 1; i >= 0; i-- {
		reversed += string(Palindrome_Data[i])
	}

	if Palindrome_Data == reversed {
		fmt.Println("PALINDROME")
		return "Its palindrome"
	} else {
		fmt.Println("NOT PALINDROME")
		return "Not palindrome"
	}
}

//Find the first repeating element in an array.

func RepeatingElement(stringdata []string) {
	Mapdata := make(map[string]int)

	for _, val := range stringdata {
		Mapdata[val]++
		if Mapdata[val] == 2 {
			fmt.Println("First repeating element:", val)
			return
		}
	}
	fmt.Println("No repeating element found.")
}
func Data() {
	data := []string{"a", "b", "c", "d", "b", "e"}
	RepeatingElement(data)
}

// Count occurrences of a target number
func Counting(array []string, target string) {
	data := make(map[string]int)
	// var data map[string]int - Uninitialized map: var data map[string]int declares a map but does not initialize it. Accessing or writing to it will cause a runtime panic.
	for _, v := range array {
		data[v]++
		if v == target {
			fmt.Println("Current count of target:", data[v])
		}

	}

	fmt.Println("Final counts:", data)
}

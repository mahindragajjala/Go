package questions

import "fmt"

/*
Bubble Sort on String Array
Sort an array of strings alphabetically using Bubble Sort.
*/

func Bubble_sort_string_array() {

	stringdata := []string{"b", "e", "a", "f", ""}
	fmt.Println("before:", stringdata)
	n := len(stringdata)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if stringdata[j] > stringdata[j+1] {
				stringdata[j], stringdata[j+1] = stringdata[j+1], stringdata[j]
			}
		}
	}
	fmt.Println("after:", stringdata)
}

/*
Start from the beginning of the array.
Compare each pair of adjacent elements.
If the first element is greater than the second (in lexicographical order), swap them.
After one full pass, the largest string (last in dictionary order) will be at the end.
Repeat the process, but each time ignore the last sorted elements (they're already in place).
Continue until the entire array is sorted.
*/

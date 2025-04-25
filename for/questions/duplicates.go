package questions

import "fmt"

/*
Write a program to remove all duplicates from an array using a for loop.
Let the array be: [1, 2, 2, 3, 4, 4, 5]

➡️ Expected Output: [1, 2, 3, 4, 5]

➡️ Goal: Use loops and condition checks to filter unique elements (hint: use an auxiliary array or set logic).
*/
func Duplicates() {
	data := []int{1, 2, 2, 3, 4, 4, 5}
	var newslice []int

	for i := 0; i < len(data); i++ {
		duplicate := false
		for j := 0; j < len(newslice); j++ {
			if data[i] == newslice[j] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			newslice = append(newslice, data[i])
		}
	}
	fmt.Println("Unique elements:", newslice)
}

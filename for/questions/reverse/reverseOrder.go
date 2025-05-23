package questions

import "fmt"

/*
Write a program to print numbers from N to 1 (reverse order).
Let N = 10.

➡️ Goal: Understand reverse looping — starting high, counting down.
*/
func ReverseOrder() {
	//#1
	N := 10
	for i := N; i >= 1; i-- {
		fmt.Println(i)
	}

	//#2
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 0             // start index
	j := len(data) - 1 // end index

	for i < j { // loop until middle
		temp := data[i]   // store left value in temp
		data[i] = data[j] // assign right value to left
		data[j] = temp    // assign temp to right
		i = i + 1         // move left pointer forward
		j = j - 1         // move right pointer backward
	}

	fmt.Println("Reversed array:", data)

	//#3
	data3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Reverse the slice
	for i, j := 0, len(data3)-1; i < j; i, j = i+1, j-1 {
		data3[i], data3[j] = data3[j], data3[i]
	}

	fmt.Println("Reversed array:", data3)
}

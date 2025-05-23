package questions

import "fmt"

// Find the missing number in the array
func Missing_number() {
	array := []int{1, 2, 4, 5, 6}
	n := len(array) + 1 // Total numbers supposed to be (1 to 6)

	expected_sum := (n * (n + 1)) / 2 // Sum of 1 to n
	actual_sum := 0

	for _, value := range array {
		actual_sum += value
	}

	missing_number := expected_sum - actual_sum
	fmt.Println(missing_number)
}

/*
Array = [1, 2, 4, 5, 6]
n = 5 + 1 = 6
Expected Sum = (6×7)/2 = 21
Actual Sum = 1 + 2 + 4 + 5 + 6 = 18
Missing Number = 21 - 18 = 3
*/

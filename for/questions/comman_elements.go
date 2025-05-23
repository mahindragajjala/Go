package questions

import "fmt"

/*
Write a program to find all common elements in two arrays using a for loop.
Let the arrays be:
A = [1, 2, 3, 4, 5]
B = [3, 4, 5, 6, 7]

➡️ Expected Output: [3, 4, 5]

➡️ Goal: Practice using loops for comparing elements across multiple data structures.
*/

func Comman_elements() {
	A := []int{1, 2, 3, 4, 5}
	B := []int{3, 4, 5, 6, 7}
	var comman_elements []int

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				// Check if the element is already added to avoid duplicates
				found := false
				for _, value := range comman_elements {
					if value == A[i] {
						found = true
						break
					}
				}
				if !found {
					comman_elements = append(comman_elements, A[i])
				}
			}
		}
	}
	fmt.Println(comman_elements)
}

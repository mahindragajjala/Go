package questions

import "fmt"

/*
Write a program to rotate an array to the right by 1 position using a for loop.
Let the array be: [1, 2, 3, 4, 5]

➡️ Expected Output: [5, 1, 2, 3, 4]

➡️ Goal: Apply loop logic to shift and rearrange array elements.
*/
func Rotate_array() {
	OLD_array := []int{1, 2, 3, 4, 5}
	NEW_array := []int{}

	for i := len(OLD_array) - 1; i <= 4; i++ {
		NEW_array = append(NEW_array, OLD_array[i])
		for j := 0; j <= len(OLD_array)-1; j++ {
			NEW_array = append(NEW_array, OLD_array[j])
		}
	}
	fmt.Println(NEW_array)
}

/*
Outer loop condition is faulty:

for i := len(OLD_array) - 1; i <= 4; i++ {
i starts at 4, and since i <= 4, the loop runs only once.

You don't need this outer loop at all if you're only rotating once.

Inner loop always appends the full array, so the result ends up being:

[5, 1, 2, 3, 4, 5] — which is one element too many.
*/

func Rotate_array_correct() {
	OLD_array := []int{1, 2, 3, 4, 5}
	NEW_array := make([]int, len(OLD_array))

	// Store last element
	last := OLD_array[len(OLD_array)-1]

	// Shift elements to the right
	for i := len(OLD_array) - 1; i > 0; i-- {
		NEW_array[i] = OLD_array[i-1]
	}

	// Put last element at first position
	NEW_array[0] = last

	fmt.Println(NEW_array)
}

/*
Step 1: Store the last element
-------------------------------
last = OLD_array[4] = 5

Step 2: Create a new array to hold rotated values
--------------------------------------------------
NEW_array = [_, _, _, _, _] (same length as OLD_array)

Step 3: Shift elements to the right (loop from end to beginning)
----------------------------------------------------------------
Loop: for i := 4; i > 0; i--
        NEW_array[i] = OLD_array[i-1]

        i = 4 → NEW_array[4] = OLD_array[3] → 4
        i = 3 → NEW_array[3] = OLD_array[2] → 3
        i = 2 → NEW_array[2] = OLD_array[1] → 2
        i = 1 → NEW_array[1] = OLD_array[0] → 1

Result after loop:
NEW_array = [_, 1, 2, 3, 4]

Step 4: Set last element at the start
-------------------------------------
NEW_array[0] = last → 5

Final Result:
-------------
NEW_array = [5, 1, 2, 3, 4]

*/

package questions

import "fmt"

// LeftRotate rotates the array to the left by 1 position.
func Rotation() {
	old_array := []int{1, 2, 3, 4, 5}
	new_array := make([]int, len(old_array))

	for i := 0; i < len(old_array)-1; i++ {
		new_array[i] = old_array[i+1]
	}
	new_array[len(old_array)-1] = old_array[0]

	fmt.Println("Old Array:", old_array)
	fmt.Println("Rotated Array:", new_array)
}

func Rotation_signal() {
	Rotation()
}

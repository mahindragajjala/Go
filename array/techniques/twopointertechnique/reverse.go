package twopointertechnique

/*
Reverse an Array In-Place
- Reverse the elements of an array without using extra space.

*/
func Reverse(array [5]int) [5]int {
	left := 0
	right := len(array) - 1
	for left < right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
	return array
}

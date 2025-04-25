package questions

import "fmt"

func Insertion_sort() {
	array := []int{5, 4, 10, 1, 6, 2}
	n := len(array)
	for i := 1; i < n; i++ {
		temp := array[i]
		j := i - 1
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = temp
	}
	fmt.Println(array)
}

/*
Start from index 1 (i = 1) to the end of the array.
Store the current value in temp.
Shift all elements greater than temp one position to the right.
Insert temp at the right location.

Real-Life Analogy:
Imagine you're picking up playing cards one by one and sorting them in your hand.
You pick the first card — you just hold it.
Then, for each next card, you look from right to left in your hand and find where it fits.
You shift cards to the right if needed and insert the new card in the correct position.


Card Game	                          Insertion Sort Code

Pick the next card                    for i := 1; i < n; i++
Hold it temporarily	                  temp := array[i]
Look left to find where it fits	      for j := i - 1; j >= 0 && array[j] > temp
Shift cards to make space	          array[j+1] = array[j]
Insert the card	                      array[j+1] = temp
*/

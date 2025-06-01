package problems

/*

	Write a function that returns the total number of elements across all rows, not just the outer slice length.

*/
func Length_Calculation() int {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	lengthdata := len(data[0])
	lengthdata3 := len(data[1])
	lengthdata4 := len(data[2])
	totalvalues := lengthdata + lengthdata3 + lengthdata4
	return totalvalues
}
func OrSumOfrows() int {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	totalValues := 0
	for _, row := range data {
		totalValues += len(row)
	}
	return totalValues
}

package problems

import "fmt"

//Add 4 new elements to the last row of the board. Print the result.

func Appending_in_inner_slice(board [][]string) {
	fmt.Println("comming-->")
	board[2] = append(board[2], "hello", "slice", "adding", "elements")
	fmt.Println(board)

}

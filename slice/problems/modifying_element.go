package problems

import "fmt"

func Modifying_Elements(Tic_tac_toe_board [][]string) string {
	/*
		Write code to update the center element of the tic-tac-toe board to 0 and print the modified board.
	*/

	//finding the row and colomn of the matrix
	row := len(Tic_tac_toe_board)
	colon := len(Tic_tac_toe_board[0])

	//find the index of the matix of middle number
	row_index := row / 2
	colon_index := colon / 2
	Tic_tac_toe_board[row_index][colon_index] = "OO"
	// fmt.Println(Tic_tac_toe_board)
	for _, value := range Tic_tac_toe_board {
		fmt.Println("After modifying:", value)
	}

	return "Modified the Tic_tac_toe_board"
}

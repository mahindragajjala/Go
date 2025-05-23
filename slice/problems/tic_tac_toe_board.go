package problems

import "fmt"

//Create a 2D slice representing a tic-tac-toe board (3x3) filled with numbers 1 to 9. Print the board.
func Tic_tac_toe_board() {
	Tic_tac_toe_board := [][]string{{"o", "x", "x"}, {"o", "o", "o"}, {"x", "x", "o"}}
	for _, value := range Tic_tac_toe_board {
		fmt.Println(value)
	}
	Modifying_Elements(Tic_tac_toe_board)
	Appending_in_inner_slice(Tic_tac_toe_board)
}

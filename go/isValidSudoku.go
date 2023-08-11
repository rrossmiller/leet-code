package main

import "fmt"

func main() {

	t := [9]string{"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79"}

	in := make([][]byte, 0)
	for i := range t {
		in = append(in, []byte(t[i]))
	}
	isValidSudoku(in)
}

func isValidSudoku(board [][]byte) bool {
	for i := range board {
		fmt.Println(i, string(board[i]))
	}
	return false
}


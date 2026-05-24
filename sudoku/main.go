package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isValid(board [9][9]int, row, col, num int) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

func solve(board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(*board, row, col, num) {
						board[row][col] = num
						if solve(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func printError() {
	for _, ch := range "Error" {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func printBoard(board [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col > 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune('0' + board[row][col]))
		}
		z01.PrintRune('\n')
	}
}

func main() {
	args := os.Args[1:]

	// Must have exactly 9 arguments
	if len(args) != 9 {
		printError()
		return
	}

	var board [9][9]int

	for i, row := range args {
		// Each row must be exactly 9 characters
		if len(row) != 9 {
			printError()
			return
		}
		for j, ch := range row {
			if ch == '.' {
				board[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				num := int(ch - '0')
				// Check if placing this number is valid before solving
				board[i][j] = 0
				if !isValid(board, i, j, num) {
					printError()
					return
				}
				board[i][j] = num
			} else {
				printError()
				return
			}
		}
	}

	if !solve(&board) {
		printError()
		return
	}

	printBoard(board)
}

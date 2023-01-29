package stateoftictactoe

import (
	"errors"
	"fmt"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	var (
		xCount int
		oCount int
		sCount int

		xWin bool
		oWin bool
	)

	count := func(line []byte, count bool) {
		var x, o int

		for _, g := range line {
			switch g {
			case 'X':
				x++
				if count {
					xCount++
				}
			case 'O':
				o++
				if count {
					oCount++
				}
			default:
				if count {
					sCount++
				}
			}
		}

		if x == 3 {
			xWin = true
		}

		if o == 3 {
			oWin = true
		}
	}

	grid := make([][]byte, 3, 3)
	for i := 0; i < 3; i++ {
		grid[i] = []byte(board[i])
	}

	for i := 0; i < 3; i++ {
		// find by row
		count(grid[i], true)
		// find by column
		count([]byte{grid[0][i], grid[1][i], grid[2][i]}, false)
	}

	// find by diagonal
	count([]byte{grid[0][0], grid[1][1], grid[2][2]}, false)
	count([]byte{grid[0][2], grid[1][1], grid[2][0]}, false)

	if !(xCount == oCount || xCount == oCount+1) {
		return "", errors.New("Board is invliad")
	}

	if xWin && oWin {
		return "", errors.New("Board is invliad, the game already ended")
	}

	if xWin || oWin {
		return Win, nil
	}

	if sCount == 0 {
		return Draw, nil
	}

	return Ongoing, nil

}

func debug(board []string) {
	for i := 0; i < 3; i++ {
		row := board[i]
		if i == 0 {
			fmt.Printf("   1 2 3\n")
			fmt.Printf(" |------\n")
		}
		fmt.Printf("%d| ", i+1)
		for j := 0; j < 3; j++ {
			fmt.Printf("%c ", row[j])
		}
		fmt.Println()
	}
	fmt.Println()
}

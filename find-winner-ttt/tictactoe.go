/*
Package t1275

Tic-tac-toe is played by two players A and B on a 3 x 3 grid. The rules of Tic-Tac-Toe are:

  - Players take turns placing characters into empty squares ' '.

  - The first player A always places 'X' characters, while the second player B always places 'O' characters.

  - 'X' and 'O' characters are always placed into empty squares, never on filled ones.

  - The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.

  - The game also ends if all squares are non-empty.

  - No more moves can be played if the game is over.

Given a 2D integer array moves where moves[i] = [rowi, coli] indicates that the ith move will be played on
grid[rowi][coli]. return the winner of the game if it exists (A or B). In case the game ends in a draw
return "Draw".

If there are still movements to play return "Pending".

You can assume that moves is valid (i.e., it follows the rules of Tic-Tac-Toe), the grid is initially empty,
and A will play first.
*/
package t1275

import (
	"fmt"
	"log"
	"reflect"
)

func isThereAWinner(board *[3][3]rune) (bool, string) {
	winnerWasFound := false
	winner := ""
	wa := [3]rune{'A', 'A', 'A'}
	wb := [3]rune{'B', 'B', 'B'}

	v1 := (*board)[0]
	v2 := (*board)[1]
	v3 := (*board)[2]
	h1 := [3]rune{(*board)[0][0], (*board)[1][0], (*board)[2][0]}
	h2 := [3]rune{(*board)[0][1], (*board)[1][1], (*board)[2][1]}
	h3 := [3]rune{(*board)[0][2], (*board)[1][2], (*board)[2][2]}
	d1 := [3]rune{(*board)[0][0], (*board)[1][1], (*board)[2][2]}
	d2 := [3]rune{(*board)[2][0], (*board)[1][1], (*board)[0][2]}

	if reflect.DeepEqual(wa, h1) ||
		reflect.DeepEqual(wa, h2) ||
		reflect.DeepEqual(wa, h3) ||
		reflect.DeepEqual(wa, d1) ||
		reflect.DeepEqual(wa, d2) ||
		reflect.DeepEqual(wa, v1) ||
		reflect.DeepEqual(wa, v2) ||
		reflect.DeepEqual(wa, v3) {
		winner = "A"
		winnerWasFound = true
	}
	if reflect.DeepEqual(wb, h1) ||
		reflect.DeepEqual(wb, h2) ||
		reflect.DeepEqual(wb, h3) ||
		reflect.DeepEqual(wb, d1) ||
		reflect.DeepEqual(wb, d2) ||
		reflect.DeepEqual(wb, v1) ||
		reflect.DeepEqual(wb, v2) ||
		reflect.DeepEqual(wb, v3) {
		winner = "B"
		winnerWasFound = true
	}
	log.Printf("H1 %v\n", h1)
	log.Printf("H2 %v\n", h2)
	log.Printf("H3 %v\n", h3)
	return winnerWasFound, winner
}

func isTheGameEnded(board *[3][3]rune) bool {
	aMoves := 0
	bMoves := 0
	for _, line := range *board {
		for _, field := range line {
			if field == 'A' {
				aMoves += 1
			}
			if field == 'B' {
				bMoves += 1
			}
		}
	}
	return aMoves+bMoves == 9
}

func analysis(board *[3][3]rune) string {
	status := "New"
	winnerWasFound, winner := isThereAWinner(board)
	if winnerWasFound {
		return winner
	} else if isTheGameEnded(board) {
		status = "Draw"
	} else {
		status = "Pending"
	}
	return status
}

func makeMove(step int, move []int, board *[3][3]rune) string {
	var marker rune
	if step%2 != 0 {
		marker = 'B'
	} else {
		marker = 'A'
	}
	if (*board)[move[0]][move[1]] != '0' {
		log.Fatal("Bad move! Field is filled.")
	}
	(*board)[move[0]][move[1]] = marker
	return analysis(board)
}

func initBoard() (string, [3][3]rune) {
	var b [3][3]rune
	for i, l := range b {
		for j := range l {
			b[i][j] = '0'
		}
	}
	return "New", b
}

func show(board *[3][3]rune) {
	fmt.Println("-----")
	for _, l := range *board {
		fmt.Printf("%c %c %c\n", l[0], l[1], l[2])
	}
	fmt.Println("-----")
}

func tictactoe(moves [][]int) string {
	result, board := initBoard()
	for i, m := range moves {
		result = makeMove(i, m, &board)
	}
	show(&board)
	log.Println("Winner: ", result)
	return result
}

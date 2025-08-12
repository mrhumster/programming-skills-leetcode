/*
Package t54

Given an m x n matrix, return all elements of the matrix in spiral order.
*/
package t54

import (
	"fmt"
	"log"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 1 {
		return matrix[0]
	}

	i := 0
	j := 0
	cols := len(matrix[0])
	rows := len(matrix)
	minrow := 0
	maxrow := rows - 1
	mincol := 0
	maxcol := cols - 1
	line := make([]int, rows*cols)
	direction := 'R'
	if len(matrix[0]) == 1 {
		direction = 'D'
	}
	for _, line := range matrix {
		for _, v := range line {
			fmt.Printf("%3d ", v)
		}
		fmt.Println()
	}
	for k := range line {
		log.Printf(
			"direct = %c, i = %d, j = %d, minrow = %d, matrix[i][j] = %d",
			direction,
			i,
			j,
			minrow,
			matrix[i][j],
		)
		line[k] = matrix[i][j]
		switch {
		case direction == 'R' && j < maxcol:
			j += 1
		case direction == 'R' && j == maxcol:
			direction = 'D'
			i += 1
			minrow += 1
		case direction == 'D' && i < maxrow:
			i += 1
		case direction == 'D' && i == maxrow:
			direction = 'L'
			j -= 1
			maxcol -= 1
		case direction == 'L' && j > mincol:
			j -= 1
		case direction == 'L' && j == mincol:
			direction = 'U'
			i -= 1
			maxrow -= 1
		case direction == 'U' && i > minrow:
			i -= 1
		case direction == 'U' && i == minrow:
			direction = 'R'
			j += 1
			mincol += 1
		}
	}
	return line
}

/*
Package t1572

# Matrix Diagonal Sum

Given a square matrix mat, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and
all the elements on the secondary diagonal that are not part of the primary diagonal.
*/
package t1572

func diagonalSum(mat [][]int) int {
	sum := 0
	if len(mat) == 1 {
		return mat[0][0]
	}
	for i, line := range mat {
		for j, v := range line {
			if i == j {
				sum += v
			}
			if i+j == len(line)-1 {
				sum += v
			}
			if i == j && i+j == len(line)-1 {
				sum -= mat[i][j]
			}
		}
	}
	return sum
}

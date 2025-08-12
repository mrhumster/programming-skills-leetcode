/*
Package t73

Given an m x n integer matrix matrix, if an element is 0,
set its entire row and column to 0's.
*/
package t73

type Pos struct {
	i int
	j int
}

func analize(matrix [][]int) []Pos {
	var zs []Pos
	for i, line := range matrix {
		for j, v := range line {
			if v == 0 {
				zs = append(zs, Pos{i, j})
				for jl := range len(line) {
					if jl != j {
						zs = append(zs, Pos{i, jl})
					}
				}
				for il := range len(matrix) {
					if il != i {
						zs = append(zs, Pos{il, j})
					}
				}
			}
		}
	}
	return zs
}

func setZeroes(matrix [][]int) [][]int {
	for _, pos := range analize(matrix) {
		matrix[pos.i][pos.j] = 0
	}
	return matrix
}


/* Package t2352
*
* Given a 0-indexed n x n integer matrix grid, return the number of
* pairs (ri, cj) such that row ri and column cj are equal.
*
* A row and column pair is considered equal if they contain the same
* elements in the same order (i.e., an equal array).
 */

package t2352

import "slices"

func equalPairs(grid [][]int) int {
	count := 0
	cols := make([][]int, 0, len(grid))
	for i := range grid {
		col := make([]int, 0, len(grid))
		for j := range grid {
			col = append(col, grid[j][i])
		}
		cols = append(cols, col)
	}
	for i := range grid {
		for j := range cols {
			if slices.Equal(grid[i], cols[j]) {
				count++
			}
		}
	}

	return count
}

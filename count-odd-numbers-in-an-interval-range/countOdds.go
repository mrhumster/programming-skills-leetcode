/*
Package t1523
Given two non-negative integers low and high.
Return the count of odd numbers between low and high (inclusive).
*/
package t1523

func countOdds(low int, high int) int {
	diff := high - low
	col := 0
	switch {
	case (diff+1)%2 == 0:
		col = (diff + 1) / 2
	case (diff+1)%2 != 0:
		col = (diff + 1) / 2
		if low%2 != 0 && high%2 != 0 {
			col += 1
		}
	}
	return col
}


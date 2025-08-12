package t1572

import "testing"

func TestDiagonalSum(t *testing.T) {
	data := []struct {
		name     string
		mat      [][]int
		expected int
	}{
		{
			"Ex.1",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			25,
		},
		{
			"Ex.2",
			[][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}},
			8,
		},
		{
			"Ex.3",
			[][]int{{5}},
			5,
		},
	}
	for _, d := range data {
		resp := diagonalSum(d.mat)
		if resp != d.expected {
			t.Errorf("⚠️ Test `%s` failed. Expected `%d`, but got `%d`.", d.name, d.expected, resp)
		}
	}
}

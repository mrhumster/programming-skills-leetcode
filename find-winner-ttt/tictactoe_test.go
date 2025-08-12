package t1275

import "testing"

func TestTictactoe(t *testing.T) {
	data := []struct {
		name     string
		moves    [][]int
		expected string
	}{
		{
			"Ex.1",
			[][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}},
			"A",
		},
		{
			"Ex.2",
			[][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
			"B",
		},
		{
			"Ex.3",
			[][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}},
			"Draw",
		},
		{
			"Ex.4",
			[][]int{{2, 0}, {1, 1}, {0, 2}, {2, 1}, {1, 2}, {1, 0}, {0, 0}, {0, 1}},
			"B",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := tictactoe(d.moves)
			if resp != d.expected {
				t.Errorf(
					"\n⚠️ Test `%s` failed. Expected `%s`, but got `%s`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

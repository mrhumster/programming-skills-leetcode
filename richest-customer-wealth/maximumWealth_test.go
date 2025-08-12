package t1672

import "testing"

func TestMaximumWealth(t *testing.T) {
	data := []struct {
		name     string
		accounts [][]int
		expected int
	}{
		{
			"Ex.1",
			[][]int{{1, 2, 3}, {3, 2, 1}},
			6,
		},
		{
			"Ex.2",
			[][]int{{1, 5}, {7, 3}, {3, 5}},
			10,
		},
		{
			"Ex.3",
			[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			17,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := maximumWealth(d.accounts)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed! Expected `%d`, but got `%d`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

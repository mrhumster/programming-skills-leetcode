package t1232

import "testing"

func TestCheckStraightLine(t *testing.T) {
	data := []struct {
		name        string
		coordinates [][]int
		expected    bool
	}{
		{
			"Ch.1",
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
			true,
		},
		{
			"Ch.2",
			[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
			false,
		},
		{
			"Ch.79",
			[][]int{{0, 0}, {0, 1}, {0, -1}},
			true,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := checkStraightLine(d.coordinates)
			if d.expected != resp {
				t.Errorf(
					"⚠️ Test `%s` failed! Expected `%t`, but got `%t`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

package t11

import "testing"

func TestMaxArea(t *testing.T) {
	data := []struct {
		name     string
		height   []int
		expected int
	}{
		{"1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"2", []int{1, 1}, 1},
		{"3", []int{2, 3, 4, 5, 18, 17, 6}, 17},
		{"4", []int{1, 2, 3, 4, 5, 25, 24, 3, 4}, 24},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := maxArea(d.height)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Tesst `%s` failed. Expected `%d`, but got `%d`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}


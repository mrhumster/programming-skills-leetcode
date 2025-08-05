package t896

import "testing"

func TestIsMonotonic(t *testing.T) {
	data := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			"Example 1",
			[]int{1, 2, 2, 3},
			true,
		},
		{
			"Example 2",
			[]int{6, 5, 4, 4},
			true,
		},
		{
			"Example 3",
			[]int{1, 3, 2},
			false,
		},
		{
			"Example 4",
			[]int{2, 2, 2, 1, 4, 5},
			false,
		},
		{
			"Example 5",
			[]int{1, 1, 0},
			true,
		},
		{
			"Example 6",
			[]int{9},
			false,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			r := isMonotonic(d.nums)
			if r != d.expected {
				t.Errorf(
					"\n⚠️ Test `%s` is failed. Expected %t, but got %t.",
					d.name,
					d.expected,
					r,
				)
			}
		})
	}
}

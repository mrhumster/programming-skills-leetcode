package t1502

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	data := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			"Example 1",
			[]int{3, 5, 1},
			true,
		},
		{
			"Example 2",
			[]int{1, 2, 4},
			false,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			r := canMakeArithmeticProgression(d.arr)
			if r != d.expected {
				t.Errorf(
					"⚠️ Test `%s` is failed. Expected %#v, but got %#v.",
					d.name,
					d.expected,
					r,
				)
			}
		})
	}
}

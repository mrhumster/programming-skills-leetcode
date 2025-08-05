package t1822

import "testing"

func TestArraySign(t *testing.T) {
	data := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			"Example 1",
			[]int{-1, -2, -3, -4, 3, 2, 1},
			1,
		},
		{
			"Example 2",
			[]int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24},
			-1,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := arraySign(d.nums)
			if d.expect != result {
				t.Errorf(
					"⚠️ Test `%s` is failed. Expected %d, but got %d", d.name, d.expect, result)
			}
		})
	}
}

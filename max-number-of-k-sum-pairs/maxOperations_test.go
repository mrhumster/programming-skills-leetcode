package t1679

import "testing"

func TestMaxOperartions(t *testing.T) {
	data := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"1", []int{1, 2, 3, 4}, 9, 2},
		{"2", []int{3, 1, 3, 4, 3}, 6, 1},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := maxOperations(d.nums, d.k)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed. Expected `%d`, but got `%d`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

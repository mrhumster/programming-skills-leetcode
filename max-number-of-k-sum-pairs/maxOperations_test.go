package t1679

import "testing"

func TestMaxOperartions(t *testing.T) {
	data := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"1", []int{1, 2, 3, 4}, 5, 2},
		{"2", []int{3, 1, 3, 4, 3}, 6, 1},
		{"3", []int{4,4,1,3,1,3,2,2,5,5,1,5,2,1,2,3,5,4}, 2, 2},
		{"4", []int{2,2,2,3,1,1,4,1}, 4, 2},
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

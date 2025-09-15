package t2352

import "testing"

func Test_eualPairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid [][]int
		want int
	}{
		// TODO: Add test cases.
		{"Example 1", [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
		{"Example 2", [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equalPairs(tt.grid)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("`%s` eualPairs() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

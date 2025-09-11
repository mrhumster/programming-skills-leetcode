package t724

import "testing"

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		// TODO: Add test cases.
		{"1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"2", []int{1, 2, 3}, -1},
		{"3", []int{2, 1, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

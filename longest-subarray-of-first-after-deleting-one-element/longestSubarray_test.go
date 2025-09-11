package t1493

import "testing"

func Test_longestSubarray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		// TODO: Add test cases.
		{"1", []int{1, 1, 0, 1}, 3},
		{"2", []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{"3", []int{1, 1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubarray(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("%s longestSubarray() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

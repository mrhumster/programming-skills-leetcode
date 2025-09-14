package t1207

import "testing"

func Test_uniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		want bool
	}{
		// TODO: Add test cases.
		{"1", []int{1, 2, 2, 1, 1, 3}, true},
		{"2", []int{1, 2}, false},
		{"3", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueOiccurrences(tt.arr)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

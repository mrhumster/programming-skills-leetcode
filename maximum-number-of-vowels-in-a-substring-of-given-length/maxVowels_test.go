package t1456

import "testing"

func Test_maxVowels(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		k    int
		want int
	}{
		// TODO: Add test cases.
		{"1","abciiidef",3, 3},
		{"Example 2","aeiou", 2, 2},
		{"Example 3","leetcode", 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxVowels(tt.s, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("%s maxVowels() = %v, want %v",tt.name, got, tt.want)
			}
		})
	}
}


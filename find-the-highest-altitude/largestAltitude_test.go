package t1732

import "testing"

func Test_largestAltitude(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		gain []int
		want int
	}{
		// TODO: Add test cases.
		{"1", []int{-5, 1, 5, 0, -7}, 1},
		{"2", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestAltitude(tt.gain)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("%s largestAltitude() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

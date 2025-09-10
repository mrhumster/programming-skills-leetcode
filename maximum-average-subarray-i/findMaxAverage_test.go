package t643

import "testing"


func Test_findMaxAverage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want float64
	}{
		// TODO: Add test cases.
		{"1", []int{1,12,-5,-6,50,3}, 4, 12.750000},
		{"2", []int{5}, 1, 5.00000},
		{"3", []int{0,4,0,3,2}, 1, 4.00000},
		{"5", []int{4,2,1,3,3}, 2, 3.00000},
		{"6", []int{4,0,4,3,3}, 5, 2.80000},
		{"7", []int{9,7,3,5,6,2,0,8,1,9}, 6, 5.333333},
	}
	for _, tt := range tests { 
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverage(tt.nums, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}


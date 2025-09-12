package t2215

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		// TODO: Add test cases.
		{"Example 1", []int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{"Example 2", []int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.nums1, tt.nums2)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("`%s` findDifference() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

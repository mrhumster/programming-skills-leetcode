package p283

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	data := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"Example 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Example 2", []int{0, 0, 1}, []int{1, 0, 0}},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := moveZeroes(d.nums)
			if !reflect.DeepEqual(result, d.expected) {
				t.Errorf(
					"⚠️ Test `%s` wrong! Expected %#v, but got %#v",
					d.name,
					d.expected,
					result,
				)
			}
		})
	}
}

package t66

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	data := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			"Example 1",
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},
		{
			"Example 2",
			[]int{9},
			[]int{1, 0},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := plusOne(d.digits)
			if !reflect.DeepEqual(d.expected, result) {
				t.Errorf(
					"⚠️ Wrong test `%s`! Expected %#v, but got %#v",
					d.name,
					d.expected,
					result,
				)
			}
		})
	}
}

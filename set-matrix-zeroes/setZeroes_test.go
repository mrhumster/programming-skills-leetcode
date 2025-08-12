package t73

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	data := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			"Ex.1",
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			"Ex.2",
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := setZeroes(d.matrix)
			if !reflect.DeepEqual(d.expected, resp) {
				t.Errorf(
					"⚠️ Test `%s` failed! Expected: \n%#v \n but got \n%#v",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

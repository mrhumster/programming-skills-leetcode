package t54

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	data := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			"Ex.1",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"Ex.2",
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			"Ex.3",
			[][]int{{1}},
			[]int{1},
		},
		{
			"Ex.15",
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
				{17, 18, 19, 20},
				{21, 22, 23, 24},
			},
			[]int{
				1,
				2,
				3,
				4,
				8,
				12,
				16,
				20,
				24,
				23,
				22,
				21,
				17,
				13,
				9,
				5,
				6,
				7,
				11,
				15,
				19,
				18,
				14,
				10,
			},
		},
		{
			"Ex.4",
			[][]int{{3}, {2}},
			[]int{3, 2},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := spiralOrder(d.matrix)
			if !reflect.DeepEqual(d.expected, resp) {
				t.Errorf(
					"⚠️ Test `%s` failed. Expected `%#v`, but got `%#v`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

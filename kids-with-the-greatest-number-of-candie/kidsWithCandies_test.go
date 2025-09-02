package t1431

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	data := []struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{"1", []int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{"2", []int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{"3", []int{12, 1, 12}, 10, []bool{true, false, true}},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := kidsWithCandies(d.candies, d.extraCandies)
			if !reflect.DeepEqual(resp, d.expected) {
				t.Errorf("⚠️ Test `%s` failed! Expected %v, but got %v.", d.name, d.expected, resp)
			}
		})
	}
}

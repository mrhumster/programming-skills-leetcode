package t238

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	data := []struct {
		name string
		nums []int
		resp []int
	}{
		{"1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"3", []int{4, 3, 2, 1, 2}, []int{12, 16, 24, 48, 24}},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := productExceptSelf(d.nums)
			if !reflect.DeepEqual(resp, d.resp) {
				t.Errorf("⚠️ Test `%s` failed! Expected `%v`, but got `%v`.", d.name, d.resp, resp)
			}
		})
	}
}

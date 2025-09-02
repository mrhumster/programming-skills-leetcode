package t605

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	data := []struct {
		name      string
		flowerbed []int
		n         int
		expected  bool
	}{
		{"1", []int{1, 0, 0, 0, 1}, 1, true},
		{"2", []int{1, 0, 0, 0, 1}, 2, false},
		{"3", []int{1, 0, 0, 0, 0, 1}, 2, false},
		{"4", []int{1, 0, 0, 0, 0, 0, 1}, 2, true},
		{"5", []int{0, 0, 1, 0, 1}, 1, true},
		{"6", []int{0}, 1, true},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := canPlaceFlowers(d.flowerbed, d.n)
			if resp != d.expected {
				t.Errorf("⚠️ Test `%s` failed. Expected %t, but got %t.", d.name, d.expected, resp)
			}
		})
	}
}


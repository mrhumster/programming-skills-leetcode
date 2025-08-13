package t860

import "testing"

func TestLemonadeChange(t *testing.T) {
	data := []struct {
		name     string
		bills    []int
		expected bool
	}{
		{"Ex.1", []int{5, 5, 5, 10, 20}, true},
		{"Ex.2", []int{5, 5, 10, 10, 20}, false},
		{"Ex.10", []int{10, 10}, false},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := lemonadeChange(d.bills)
			if resp != d.expected {
				t.Errorf(
					"\n⚠️ Test `%s` failed! Expected `%t`, but got `%t`!",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

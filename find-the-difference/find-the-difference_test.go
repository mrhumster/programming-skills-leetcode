package t389

import "testing"

func TestOfFindTheDifference(t *testing.T) {
	data := []struct {
		name     string
		s        string
		t        string
		expected byte
	}{
		{"Example 1", "abcd", "abcde", byte('e')},
		{"Example 2", "", "y", byte('y')},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := findTheDifference(d.s, d.t)
			if result != d.expected {
				t.Errorf("Expected %c, got %c", d.expected, result)
			}
		})
	}
}

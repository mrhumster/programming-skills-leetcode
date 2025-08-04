package t28

import "testing"

func TestStrStr(t *testing.T) {
	data := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{"Example 1", "sadbutsad", "sad", 0},
		{"Example 2", "leetcode", "leeto", -1},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := strStr(d.haystack, d.needle)
			if result != d.expected {
				t.Errorf("Expected %d, but %d", d.expected, result)
			}
		})
	}
}

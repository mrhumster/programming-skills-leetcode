package t242

import "testing"

func TestIsAnagram(t *testing.T) {
	data := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"Example 1", "anagram", "nagaram", true},
		{"Example 2", "rat", "car", false},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := isAnagram(d.s, d.t)
			if result != d.expected {
				t.Errorf("Expected %t, but got %t", d.expected, result)
			}
		})
	}
}

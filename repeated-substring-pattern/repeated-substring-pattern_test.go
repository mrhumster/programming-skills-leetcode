package t459

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	data := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Example 1", "abab", true},
		{"Example 2", "aba", false},
		{"Example 3", "abcabcabcabc", true},
		{"Example 4", "ababba", false},
		{"Example 5", "bb", true},
		{"Example 6", "abaababaab", true},
		{"Example 7", "aabaaba", false},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := repeatedSubstringPattern(d.s)
			if result != d.expected {
				t.Errorf(
					"\n⚠️ Test `%s` wrong. The param was like that: `%s` Expected %t, but got %t",
					d.name,
					d.s,
					d.expected,
					result,
				)
			}
		})
	}
}

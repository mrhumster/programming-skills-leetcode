package T1768

import (
	"strings"
	"testing"
)

func TestOfMergeAlternaty(t *testing.T) {
	data := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{"Example 1", "abc", "pqr", "apbqcr"},
		{"Example 2", "ab", "pqrs", "apbqrs"},
		{"Example 3", "abcd", "pq", "apbqcd"},
		{
			"Validation 1",
			"",
			"abc",
			"",
		},
		{
			"Validation 2",
			"abc",
			"",
			"",
		},
		{
			"Validation 3",
			strings.Repeat(".", 101),
			"pqr",
			"",
		},
		{
			"Validation 4",
			"abc",
			strings.Repeat(".", 101),
			"",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := mergeAlternately(d.word1, d.word2)
			if result != d.expected {
				t.Errorf("Expected %s, got %s", d.expected, result)
			}
		})
	}
}

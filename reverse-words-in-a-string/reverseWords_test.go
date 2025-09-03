package t151

import "testing"

func TestReverseWord(t *testing.T) {
	data := []struct {
		name     string
		s        string
		expected string
	}{
		{"1", "the sky is blue", "blue is sky the"},
		{"2", "  hello world ", "world hello"},
		{"3", "a good   example", "example good a"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := reverseWords(d.s)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed. Expected `%v`, but go `%v`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

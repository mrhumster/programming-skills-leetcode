package t345

import "testing"

func TestTeverseTowel(t *testing.T) {
	data := []struct {
		name     string
		s        string
		expected string
	}{
		{"1", "IceCreAm", "AceCreIm"},
		{"2", "leetcode", "leotcede"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := reverseVowels(d.s)
			if resp != d.expected {
				t.Errorf("⚠️ Test `%s` failed. Expected %v, but go %v.", d.name, d.expected, resp)
			}
		})
	}
}

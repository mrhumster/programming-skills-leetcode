package t392

import "testing"

func TestIsSusequence(t *testing.T) {
	data := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"1", "abc", "ahbgdc", true},
		{"2", "axc", "ahbgdc", false},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := isSubsequence(d.s, d.t)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed. Expected `%v`, but got `%v`",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}


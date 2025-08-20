package t67

import "testing"

func TestAddBinary(t *testing.T) {
	data := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{"Ex.1", "11", "1", "100"},
		{"Ex.2", "1010", "1011", "10101"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := addBinary(d.a, d.b)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed. Expected `%s`, but got `%s`!",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

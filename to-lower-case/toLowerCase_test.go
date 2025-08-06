package t709

import "testing"

func TestToLowerCase(t *testing.T) {
	data := []struct {
		name     string
		s        string
		expected string
	}{
		{
			"Ex.1",
			"Hello",
			"hello",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := toLowerCase(d.s)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` is failed. Expected %s, but go %s.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

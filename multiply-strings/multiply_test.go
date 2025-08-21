package t43

import "testing"

func TestMultiply(t *testing.T) {
	data := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{
		{
			"Ch.1",
			"2",
			"3",
			"6",
		},
		{
			"Ch.2",
			"123",
			"456",
			"56088",
		},
		{
			"Ch.3",
			"9133",
			"0",
			"0",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := multiply(d.num1, d.num2)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed! Expected `%s`, but got `%s`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

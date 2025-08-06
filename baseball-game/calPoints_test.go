package baseballgame

import "testing"

func TestCalPoints(t *testing.T) {
	data := []struct {
		name       string
		operations []string
		expected   int
	}{
		{
			"Ex.1",
			[]string{"5", "2", "C", "D", "+"},
			30,
		},
		{
			"Ex.2",
			[]string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			27,
		},
		{
			"Ex.3",
			[]string{"1", "C"},
			0,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := calPoints(d.operations)
			if resp != d.expected {
				t.Errorf("⚠️ Test `%s` failed. Expected  %d, but got %d.", d.name, d.expected, resp)
			}
		})
	}
}

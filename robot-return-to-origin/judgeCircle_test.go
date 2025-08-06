package robotreturntoorigin

import "testing"

func TestJudjeCircle(t *testing.T) {
	data := []struct {
		name     string
		moves    string
		expected bool
	}{
		{
			"Ex.1",
			"UD",
			true,
		},
		{
			"Ex.2",
			"LL",
			false,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := judgeCircle(d.moves)
			if resp != d.expected {
				t.Errorf(
					"\n⚠️ Test `%s` failed! Expected %t, but got %t.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

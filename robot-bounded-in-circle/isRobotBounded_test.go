package t1041

import "testing"

func TestIsRobotBounded(t *testing.T) {
	data := []struct {
		name         string
		instructions string
		expected     bool
	}{
		{
			"Ex.1",
			"GGLLGG",
			true,
		},
		{
			"Ex.2",
			"GG",
			false,
		},
		{
			"Ex.3",
			"GL",
			true,
		},
		{
			"Ex.108",
			"GLGLGGLGL",
			false,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := isRobotBounded(d.instructions)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed. Expected `%t`, but got `%t`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

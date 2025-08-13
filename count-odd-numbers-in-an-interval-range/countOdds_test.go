package t1523

import "testing"

func TestCountOdds(t *testing.T) {
	data := []struct {
		name     string
		low      int
		high     int
		expected int
	}{
		{
			"Ex.1",
			3,
			7,
			3,
		},
		{
			"Ex.2",
			8,
			10,
			1,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := countOdds(d.low, d.high)
			if resp != d.expected {
				t.Errorf(
					"\n⚠️ Test `%s` failed. Expected `%d`, but got `%d`!",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

